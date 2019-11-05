using System;
using System.Collections.Generic;
using System.Linq;
using System.Reflection;
using System.Text;

namespace GameBox.Service
{
    public class Entity
    {
        private Dictionary<string, Dictionary<object, MethodInfo>> rpcHandlerDict;

        internal ulong id;
        internal string type;
        internal Action<string, object[]> rpcAction;

        public ulong Id
        {
            get { return id; }
        }

        public string Type
        {
            get { return type; }
        }

        public Entity()
        {
            rpcHandlerDict = new Dictionary<string, Dictionary<object, MethodInfo>>();
            RegisterRpcHandler(this);
        }

        public void Rpc(string methodname, params object[] args)
        {
            if(rpcAction != null)
            {
                rpcAction(methodname, args);
            }
            else
            {
                throw new Exception("only user entity can send rpc msg!");
            }
        }

        internal void OnRpc(string methodname, byte[] data)
        {
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
                            if (!MsgHelper.UnpackMethod(info.Value, data, out args))
                            {
                                throw new Exception("Parameter cannot be unpacked for rpc method " + methodname + " in entity id->" + id);
                            }
                        }
                        catch (Exception e)
                        {
                            throw new Exception("OnRpc UnpackMethod Execption for method name " + methodname + " in entity id->" + id + ". exception:" + e);
                        }
                    }
                    info.Value.Invoke(info.Key, args);
                }
            }
            else
            {
                throw new Exception("rpc method name->" + methodname + " not exists in entity id->" + id);
            }
        }

        private void RegisterRpcHandler(object handler)
        {
            Dictionary<object, MethodInfo> rpcinfo = null;

            Type type = handler.GetType();
            var methods = type.GetMethods(BindingFlags.Instance | BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.Static)
                .Where(m => m.GetCustomAttributes(typeof(GBoxRPCAttribute), true).Length > 0).ToArray();

            for (int i = 0; i < methods.Length; i++)
            {
                var name = methods[i].Name;

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
    }
}
