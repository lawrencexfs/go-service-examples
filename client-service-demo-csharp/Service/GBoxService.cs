/*
* This file is part of the GameBox package.
*
* (c) Giant - MouGuangYi<mouguangyi@ztgame.com> , tanxiaoliang<tanxiaoliang@ztgame.com>
*
* For the full copyright and license information, please view the LICENSE
* file that was distributed with this source code.
*
* Document: http://192.168.150.238/GameBox/help/ 
*/

using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Threading;
using GameBox.Channel;
using GameBox.Network;
using GameBox.Service.msg;

namespace GameBox.Service
{
    /// <summary>
    /// Service网络逻辑
    /// </summary>
    public abstract class GBoxService
    {
        public long Rtt
        {
            get;
            private set;
        }

        protected event Action WhenConnected;
        protected event Action<Exception> WhenClosed;
        protected event Action<Exception> WhenError;

        private GaiaPacker packer;

        private Events events;

        private IMsgDefined msgDefined;
        //private PackerPolicy packerPolicy;

        private List<object> recvMsgList;
        private List<object> tmpMsgList;
        //private readonly byte[] sentBuffer;
        //private MemoryStream protoStream;
        //private MemoryStream protoRecvStream;
        //private MMStream binaryStream;
        //private MMStream binaryRecvStream;
        //private const int sentBufferMaxCount = 8192;
        //private const int headSize = 6;

        private INetwork network;
        private bool bConnected;

        public ulong entity_id; //当前玩家实体id
        public string entity_type; //当前玩家实体类型
        private Dictionary<ulong, Entity> entityDict;
        private Dictionary<string, Type> entityTypeDict;
        public Action<Entity> OnEntityCreated;
        public Action<Entity> OnEntityDestroyed;

        public Entity Player()
        {
            if (entityDict.ContainsKey(entity_id))
            {
                return entityDict[entity_id];
            }

            return null;
        }

        public GBoxService()
        {
            events = new Events();
            msgDefined = new GaiaMsgDefined();

            entityDict = new Dictionary<ulong, Entity>();
            entityTypeDict = new Dictionary<string, Type>();

            var policy = new PackerPolicy();
            policy.AddPackedPolicy(BinaryPacker.Packed);
            policy.AddPackedPolicy(ProtobufPacker.Packed);
            policy.AddUnpackedPolicy(BinaryPacker.Unpacked);
            policy.AddUnpackedPolicy(ProtobufPacker.Unpacked);

            packer = new GaiaPacker(policy, msgDefined);

            //var type = typeof(msg.Ping);
            //var msgname = type.FullName;
            //msgDefined.AddMsgDefined(msg.MsgId.PingMsgID, msgname, type);
            RegisterMsg<msg.Ping>(MsgId.PingMsgID);

            //type = typeof(Pong);
            //msgname = type.FullName;
            //msgDefined.AddMsgDefined(msg.MsgId.PongMsgID, msgname, type);
            RegisterMsg<Pong>(MsgId.PongMsgID);

            //type = typeof(RpcMsg);
            //msgname = type.FullName;
            //msgDefined.AddMsgDefined(msg.MsgId.CallMsgID, msgname, type);
            RegisterMsg<RpcMsg>(MsgId.CallMsgID);

            //type = typeof(CreateEntityNotify);
            //msgname = type.FullName;
            //msgDefined.AddMsgDefined(msg.MsgId.CreateEntityNotifyMsgID, msgname, type);
            RegisterMsg<CreateEntityNotify>(MsgId.CreateEntityNotifyMsgID, OnCreateEntity);

            RegisterRpcHandler(this);
        }

        //private void OnPong(Pong pong)
        //{
        //    lock (ping)
        //    {
        //        pingFlag = false;
        //        pingTime = 0;

        //        long nowtimestamp = (long)DateTime.UtcNow.Subtract(DateTime.MinValue).Ticks / 10000;
        //        Rtt = nowtimestamp - sendTimeStamp;
        //    }
        //}

        private void OnCreateEntity(CreateEntityNotify notify)
        {
            if (!entityTypeDict.ContainsKey(notify.EntityType))
            {
                throw new Exception("entity define name->" + notify.EntityType + " not exists!");
            }

            var newentity = Activator.CreateInstance(entityTypeDict[notify.EntityType]) as Entity;
            newentity.id = notify.EntityID;
            newentity.type = notify.EntityType;

            entityDict[notify.EntityID] = newentity;

            if (entity_type == notify.EntityType)
            {
                entity_id = notify.EntityID;
                newentity.rpcAction = Rpc;
            }

            if (OnEntityCreated != null)
            {
                OnEntityCreated(newentity);
            }
        }

        public void RegisterEntity<T>(bool isuser = false)
        {
            var type = typeof(T);
            var name = type.Name;

            if (entityTypeDict.ContainsKey(name))
            {
                throw new Exception("entity define name->" + name + " already exists!");
            }

            entityTypeDict[name] = type;

            if (isuser)
            {
                entity_type = name;
            }
        }

        public void Disconnect()
        {
            if (network != null)
            {
                StopHeart();
                
                network.Disconnect();

                network.OnPacket -= OnPacket;
                network.OnConnected -= Connected;
                network.OnClosed -= Closed;
                network.OnError -= Error;

                network = null;

                //recvMsgList = null;
                //tmpMsgList = null;
                bConnected = false;
                entityDict.Clear();
            }
        }

        public void Connect(string net, string ip, int port)
        {
            //Disconnect();
            if (network != null)
            {
                StopHeart();

                network.OnPacket -= OnPacket;
                network.OnConnected -= Connected;
                network.OnClosed -= Closed;
                network.OnError -= Error;

                network.Disconnect();

                network = null;

                //recvMsgList = null;
                //tmpMsgList = null;
                bConnected = false;
                entityDict.Clear();
            }

            try
            {
                if (net == "tcp")
                    network = GBox.Make<INetworkManager>().Create("cratos", "cratos.tcp://" + ip + ":" + port);
                else if (net == "kcp")
                    network = GBox.Make<INetworkManager>().Create("cratos", "cratos.kcp://" + ip + ":" + port + "?datashards=3&parityshards=2");
                network.OnConnected += Connected;
                network.OnClosed += Closed;
                network.OnError += Error;

                tmpMsgList = new List<object>();
                recvMsgList = new List<object>();
                eventList = new List<EventData>();

                network.OnPacket += OnPacket;
                network.SetPlugins(packer);
                network.Connect();
            }
            catch(Exception e)
            {
                if (WhenError != null)
                    WhenError.Invoke(e);
            }
        }

        enum EventType
        {
            Connected,
            Closed,
            Error
        }

        struct EventData
        {
            public EventType eventType;
            public Exception ex;
        }

        List<EventData> eventList;

        private void Connected(INetwork network, IChannel channel)
        {
            bConnected = true;
            //if (OnConnected != null)
            //    OnConnected.Invoke();

            //heartThread = new Thread(HeartThread);
            //heartThread.IsBackground = true;
            //bHeartThreadQuit = false;
            //heartThread.Start();
            lock (eventList)
            {
                eventList.Add(new EventData() { eventType = EventType.Connected, ex = null });
            }
        }

        private void Closed(INetwork network, IChannel channel, Exception ex)
        {
            bConnected = false;
            //if (OnClosed != null)
            //    OnClosed.Invoke(ex);

            //StopHeart();
            lock (eventList)
            {
                eventList.Add(new EventData() { eventType = EventType.Closed, ex = null });
            }
        }

        private void StopHeart()
        {
            if (heartThread != null)
            {
                bHeartThreadQuit = true;
                heartThread = null;
            }
        }

        private void Error(INetwork network, IChannel channel, Exception ex)
        {
            //if (OnError != null)
            //    OnError.Invoke(ex);
            lock (eventList)
            {
                eventList.Add(new EventData() { eventType = EventType.Error, ex = ex });
            }
        }

        //private long accTime = 0;
        //private long lastTime = 0;
        private long pingTime = 0;
        private bool pingFlag = false;
        private long sendTimeStamp;
        private msg.Ping ping = new msg.Ping();

        private Thread heartThread;
        private bool bHeartThreadQuit;

        //private void OnPong(Pong pong)
        //{
        //    pingFlag = false;
        //}

        private void HeartThread()
        {
            while (true)
            {
                if (bHeartThreadQuit)
                    break;

                lock (ping)
                {
                    if (!pingFlag)
                    {
                        //now send ping always to keep alive, and don't check if recv pong.
                        //Debug.Log("sending ping to server..." + this.GetType().Name);
                        sendTimeStamp = (long)DateTime.UtcNow.Subtract(DateTime.MinValue).Ticks / 10000;
                        Send(ping);
                        pingTime = 0;
                        pingFlag = true;
                    }
                    else if (pingTime >= 10000)
                    {
                        //Debug.LogError("ping time out..." + this.GetType().Name);
                        //ping time out
                        Disconnect();
                        //lock (eventList)
                        //{
                        //    eventList.Add(new EventData() { eventType = EventType.Error, ex = new TimeoutException("ping time out") });
                        //}

                        pingFlag = false;
                        pingTime = 0;
                        return;
                    }
                    else
                    {
                        //Send(ping);
                        pingTime += 1000;
                        Rtt = pingTime;
                    }
                }

                Thread.Sleep(1000);
            }
        }

        private void OnPacket(INetwork network, IChannel channel, object packet)
        {
            if (packet.GetType().Name == "Pong")
            {
                //Debug.Log("recv pong from server..." + this.GetType().Name);
                lock (ping)
                {
                    pingFlag = false;
                    pingTime = 0;

                    long nowtimestamp = (long)DateTime.UtcNow.Subtract(DateTime.MinValue).Ticks / 10000;
                    Rtt = nowtimestamp - sendTimeStamp;
                }

                return;
            }

            lock (recvMsgList)
            {
                recvMsgList.Add(packet);
            }
            //var msgInfo = MsgService.Instance.GetMsgByName(packet.GetType().Name);
            //if (msgInfo != null)
            //{
            //    Events.Fire(msgInfo.Name, packet);
            //    if (packet is IMsg)
            //    {
            //        ((IMsg)packet).Recycle();
            //    }
            //}
        }

        //private uint bkdr_hash(string msgname)
        //{
        //    uint seed = 131;
        //    uint hash = 0;

        //    for(int i = 0; i < msgname.Length; i++)
        //    {
        //        hash = hash * seed + (uint)msgname[i];
        //    }

        //    return hash;
        //}

        public void Send(object req)
        {
            if(network != null)
            {
                var result = network.SendTo(req);
                if(result != Socket.SendResults.Success && result != Socket.SendResults.Pending)
                {
                    //lock (eventList)
                    //{
                    //    eventList.Add(new EventData() { eventType = EventType.Error, ex = new TimeoutException("ping time out") });
                    //}
                }
            }
        }

        protected void RegisterMsg<T>(ushort msgid, Action<T> handler = null)
        {
            var type = typeof(T);
            var msgname = type.FullName;
            msgDefined.AddMsgDefined((ushort)msgid, msgname, type);

            if (handler != null)
                events.Add<T>(msgname, handler);
        }

        protected void UnRegisterMsg<T>(uint msgid)
        {
            var msgname = typeof(T).FullName;
            events.Remove(msgname);
            msgDefined.RemoveMsgDefined(msgid);
        }

        public virtual void Update(float delta)
        {
            if (recvMsgList == null)
                return;

            lock (recvMsgList)
            {
                tmpMsgList.AddRange(recvMsgList);
                recvMsgList.Clear();
            }

            if (tmpMsgList.Count > 0)
            {
                for (int i = 0; i < tmpMsgList.Count; i++)
                {
                    var msg = tmpMsgList[i];
                    var fullname = msg.GetType().FullName;

                    if (msg.GetType().Name == "RpcMsg")
                    {
                        OnRpc((RpcMsg)msg);
                    }
                    else
                    {
                        var msgInfo = msgDefined.GetMsgByName(fullname);
                        if (msgInfo != null)
                        {
                            events.Fire(msgInfo.Name, msg);
                            if (msg is IMsg)
                            {
                                ((IMsg)msg).Recycle();
                            }
                        }
                    }
                }

                tmpMsgList.Clear();
            }

            lock (eventList)
            {
                for (int i = 0; i < eventList.Count; i++)
                {
                    var e = eventList[i];
                    switch (e.eventType)
                    {
                        case EventType.Connected:
                            if (WhenConnected != null)
                                WhenConnected.Invoke();

                            heartThread = new Thread(HeartThread);
                            heartThread.IsBackground = true;
                            bHeartThreadQuit = false;
                            heartThread.Start();
                            break;
                        case EventType.Closed:
                            StopHeart();
                            if (WhenClosed != null)
                                WhenClosed.Invoke(null);
                            break;
                        case EventType.Error:
                            if (WhenError != null)
                                WhenError.Invoke(e.ex);
                            break;
                    }
                }

                eventList.Clear();
            }
        }

        //protected abstract void fireOnError(string errdesc);

#region rpc
        //rpc functions
        private Dictionary<string, Dictionary<object, MethodInfo>> rpcHandlerDict = new Dictionary<string, Dictionary<object, MethodInfo>>();

        public void Rpc(string methodname, params object[] args)
        {
            byte[] bytes;
            if (!MsgHelper.Pack(args, out bytes))
            {
                throw new Exception("Parameter cannot be pack for rpc call method:" + methodname);
            }

            RpcMsg req = new RpcMsg();
            //req.rpcType = rpctype;
            req.methodeName = methodname;
            req.data = bytes;

            try
            {
                Send(req);
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }

        private void OnRpc(RpcMsg msg)
        {
            var methodname = msg.methodeName;
            //for (var k = 0; k < args.Length; k++)
            //{
            //    name += "_" + args[k].GetType().Name;
            //}
            //fire(msg.methodeName, args);
            Dictionary<object, MethodInfo> rpcinfo = new Dictionary<object, MethodInfo>();
            if (rpcHandlerDict.TryGetValue(methodname, out rpcinfo))
            {
                object[] args = null;
                var iter = rpcinfo.GetEnumerator();
                bool flag = true;

                while (iter.MoveNext())
                {
                    var info = iter.Current;
                    if (flag)
                    {
                        flag = false;
                        try
                        {
                            if (!MsgHelper.UnpackMethod(info.Value, msg.data, out args))
                            {
                                throw new Exception("Parameter cannot be unpacked for rpc method " + methodname);
                            }
                        }
                        catch (Exception e)
                        {
                            throw new Exception("OnRpc UnpackMethod Execption for method name " + methodname + " exception:" + e);
                        }
                    }
                    info.Value.Invoke(info.Key, args);
                }

                return;
            }

            if(!entityDict.ContainsKey(msg.EntityID))
            {
                throw new Exception("entity id->" + msg.EntityID + " not exists!");
            }

            entityDict[msg.EntityID].OnRpc(methodname, msg.data);
        }

        protected void RegisterRpcHandler(object handler)
        {
            Dictionary<object, MethodInfo> rpcinfo = null;

            Type type = handler.GetType();
            var methods = type.GetMethods(BindingFlags.Instance | BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.Static)
                .Where(m => m.GetCustomAttributes(typeof(GBoxRPCAttribute), true).Length > 0).ToArray();
            for (int i = 0; i < methods.Length; i++)
            {
                var name = methods[i].Name;
                //var paramarr = methods[i].GetParameters();
                //for (var k = 0; k < paramarr.Length; k++)
                //{
                //    name += "_" + paramarr[k].ParameterType.Name;
                //}

                if (!rpcHandlerDict.TryGetValue(name, out rpcinfo))
                {
                    rpcinfo = new Dictionary<object, MethodInfo>();
                    rpcHandlerDict.Add(name, rpcinfo);
                }

                if (!rpcinfo.ContainsKey(handler))
                {
                    rpcinfo.Add(handler, methods[i]);
                }
                else
                {
                    throw new Exception("method " + name + " with handler " + handler.ToString() + " already exists");
                }
            }
        }

        protected void UnRegisterRpcHandler(object target)
        {
            Dictionary<object, MethodInfo> rpcinfo = new Dictionary<object, MethodInfo>();
            Type type = target.GetType();
            var methods = type.GetMethods(BindingFlags.Instance | BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.Static)
                .Where(m => m.GetCustomAttributes(typeof(GBoxRPCAttribute), true).Length > 0).ToArray();
            for (int i = 0; i < methods.Length; i++)
            {
                var name = methods[i].Name;
                //var paramarr = methods[i].GetParameters();
                //for (var k = 0; k < paramarr.Length; k++)
                //{
                //    name += "_" + paramarr[k].ParameterType.Name;
                //}

                if (rpcHandlerDict.TryGetValue(name, out rpcinfo))
                {
                    if (rpcinfo.ContainsKey(target))
                    {
                        rpcinfo.Remove(target);
                    }
                    if (rpcinfo.Count == 0)
                    {
                        rpcHandlerDict.Remove(name);
                    }
                }
            }
        }
#endregion
    }
}