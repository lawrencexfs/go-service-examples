using GameBox;
using GameBox.Channel;
using GameBox.Network;
using GameBox.Pioneer;
using GameBox.Socket;
using System;
using System.Collections.Generic;
using System.IO;
using UnityEngine;
using usercmd;

namespace Kongfu
{
    public sealed class RoomSystem : GameBox.Pioneer.System, IRoomManager, IPacker
    {
        public override void OnInit(IEntityContainer container)
        {
            var channelManager = GBox.Make<IChannelManager>();
            channelManager.Extend("kongfu", (nsp) =>
            {
                var socket = GBox.Make<ISocketManager>().Create(nsp);
                return new Channel(socket, new Fragment());
            });
        }

        public override void OnUpdate(IEntityContainer container, float deltaTime)
        {
            lock (this.packets)
            {
                while (this.packets.Count > 0)
                {
                    var packet = this.packets.Dequeue();
                    Dispatch(packet.Id, packet.Proto);
                }
            }
        }

        public void Enter(string ip, int port)
        {
            if (null != this.network)
            {
                this.network.Disconnect();
                this.network = null;
            }

            var nsp = "kongfu.tcp://" + ip + ":" + port;
            this.network = GBox.Make<INetworkManager>().Create("kongfu.game", nsp);
            this.network.OnConnected += OnConnected;
            this.network.OnClosed += OnClosed;
            this.network.OnError += OnError;
            this.network.OnPacket += OnReceive;
            this.network.SetPlugins(this);
            this.network.Connect();
        }

        public void Exit()
        {
            this.network.Disconnect();
            this.actions.Clear();
        }

        public void SendPacket(int id, object obj = null)
        {
            var pack = new MsgPacket { Id = id, Proto = (null == obj ? null : new Proto(obj)) };
            this.network.SendTo(pack);
        }

        public void OnPacket(int id, Action<IProto> action)
        {
            LazySet<Action<IProto>> handlers = null;
            if (!this.actions.TryGetValue(id, out handlers))
            {
                this.actions.Add(id, handlers = new LazySet<Action<IProto>>());
            }

            if (!handlers.Contains(action))
            {
                handlers.Add(action);
            }
        }

        public void OffPacket(int id, Action<IProto> action)
        {
            LazySet<Action<IProto>> handlers = null;
            if (this.actions.TryGetValue(id, out handlers))
            {
                handlers.Remove(action);
            }
        }

        #region IPacker
        public ArraySegment<byte> Pack(object packet)
        {
            var obj = (MsgPacket)packet;
            var bytes = PackMessage(obj.Id, null != obj.Proto ? obj.Proto.ToBytes() : null); ;
            return new ArraySegment<byte>(bytes);
        }

        public object Unpack(ArraySegment<byte> packet)
        {
            var len1 = packet.Array[packet.Offset + 0];
            var len2 = packet.Array[packet.Offset + 1];
            var len3 = packet.Array[packet.Offset + 2];
            var compressed = BitConverter.ToBoolean(packet.Array, packet.Offset + 3);
            var bodyLength = len1 | (len2 << 8) | (len3 << 16);
            var cmdId = (int)BitConverter.ToUInt16(packet.Array, packet.Offset + 4);
            byte[] bodyBytes = null;
            if (bodyLength > Fragment.CmdIdLength)
            {
                bodyBytes = new byte[bodyLength - Fragment.CmdIdLength];
                Array.Copy(packet.Array, 4 + Fragment.CmdIdLength, bodyBytes, 0, bodyLength - Fragment.CmdIdLength);
                // TODO: Handle compress flag
            }

            return new MsgPacket { Id = cmdId, Proto = (null != bodyBytes ? new Proto(bodyBytes) : null) };
        }
        #endregion

        private void OnConnected(INetwork network, IChannel channel)
        {
            Debug.Log("Connect server successfully.");

            var msg = new MsgLogin();
            msg.name = "玩家" + new System.Random().Next(1000);
            SendPacket((int)MsgTypeCmd.Login, msg);
        }

        private void OnClosed(INetwork network, IChannel channel, Exception e)
        {
            Debug.Log("Disconnect server.");
        }

        private void OnError(INetwork network, IChannel channel, Exception e)
        {
            Debug.Log("Error: " + e.Message);
        }

        private void OnReceive(INetwork network, IChannel channel, object packet)
        {
            lock (this.packets)
            {
                this.packets.Enqueue((MsgPacket)packet);
            }
        }

        private byte[] PackMessage(int type, byte[] body)
        {
            using (var stream = new MemoryStream())
            {
                using (var writer = new BinaryWriter(stream))
                {
                    uint packBodyLen = Fragment.CmdIdLength;
                    if (body != null)
                    {
                        packBodyLen += (uint)body.Length;
                    }

                    var len1 = (byte)packBodyLen;
                    var len2 = (byte)(packBodyLen >> 8);
                    var len3 = (byte)(packBodyLen >> 16);
                    writer.Write(len1);
                    writer.Write(len2);
                    writer.Write(len3);
                    writer.Write((byte)0);
                    //
                    writer.Write((ushort)type);
                    if (body != null)
                    {
                        writer.Write(body);
                    }
                }

                return stream.ToArray();
            }
        }

        private void Dispatch(int id, IProto proto)
        {
            LazySet<Action<IProto>> handlers = null;
            if (this.actions.TryGetValue(id, out handlers) && handlers.Count > 0)
            {
                foreach (var action in handlers)
                {
                    action(proto);
                }
            }
        }

        private struct MsgPacket
        {
            public int Id;
            public IProto Proto;
        }

        private INetwork network = null;
        private Dictionary<int, LazySet<Action<IProto>>> actions = new Dictionary<int, LazySet<Action<IProto>>>();
        private Queue<MsgPacket> packets = new Queue<MsgPacket>();
    }
}
