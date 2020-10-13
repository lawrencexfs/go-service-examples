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
using System.Collections;
using System.Collections.Generic;
using System.Reflection;
//using UnityEngine;

namespace GameBox.Service
{
    /// <summary>
    /// Raw Proto
    /// </summary>
    internal class RawProto
    {
        public byte[] rawData;
        public string name;

        private object unpackedData;

        public T unpack<T>()
        {
            // todo: 这里需要接入统一的打包策略
            if (unpackedData != null)
                return (T)unpackedData;
            var stream = new System.IO.MemoryStream(rawData);
            stream.SetLength(rawData.Length);
            unpackedData = (T)ProtoBuf.Meta.RuntimeTypeModel.Default.Deserialize(stream, null, typeof(T));
            return (T)unpackedData;
        }
    }

    /// <summary>
    /// 消息帮助库
    /// </summary>
    internal static class MsgHelper
    {
        public static bool Pack(object[] args, out byte[] msgBuf)
        {
            var stream = MMStreamPool.Get(Header.MaxPacketSize);
            stream.WPos = 0;
            stream.RPos = 0;

            for (var i = 0; i < args.Length; i++)
            {
                doPack(stream, args[i].GetType(), args[i]);
            }

            msgBuf = new byte[stream.WPos];
            Buffer.BlockCopy(stream.Buf, 0, msgBuf, 0, stream.WPos);
            MMStreamPool.Put(stream);
            return true;
        }

        public static void PackParam(object obj, out byte[] buf)
        {
            var stream = MMStreamPool.Get(Header.MaxPacketSize);
            stream.WPos = 0;
            stream.RPos = 0;

            doPack(stream, obj.GetType(), obj);

            buf = new byte[stream.WPos];
            Buffer.BlockCopy(stream.Buf, 0, buf, 0, stream.WPos);
            MMStreamPool.Put(stream);
        }

        private static void doPack(MMStream stream, Type type, object arg)
        {
            //if (arg == null)
            //    return;
            //Debug.Log("doPack arg.GetType():" + arg.GetType() + " arg.GetType().IsValueType:" + arg.GetType().IsValueType + " arg.GetType().IsArray:" + arg.GetType().IsArray);
            if (type == typeof(byte))
            {
                stream.WriteByte((byte)arg);
            }
            else if (type == typeof(sbyte))
            {
                stream.WriteByte((byte)arg);
            }
            else if (type == typeof(ushort))
            {
                stream.WriteUInt16((ushort)arg);
            }
            else if (type == typeof(uint))
            {
                stream.WriteUInt32((uint)arg);
            }
            else if (type == typeof(ulong))
            {
                stream.WriteUInt64((ulong)arg);
            }
            else if (type == typeof(short))
            {
                stream.WriteInt16((short)arg);
            }
            else if (type == typeof(int))
            {
                stream.WriteInt32((int)arg);
            }
            else if (type == typeof(long))
            {
                stream.WriteInt64((long)arg);
            }
            else if (type == typeof(float))
            {
                stream.WriteFloat((float)arg);
            }
            else if (type == typeof(double))
            {
                stream.WriteDouble((double)arg);
            }
            else if (type == typeof(string))
            {
                stream.WriteString((string)arg);
            }
            else if (type == typeof(byte[]))
            {
                stream.WriteBytes((byte[])arg);
            }
            else if (type == typeof(sbyte[]))
            {
                var buff = new byte[((sbyte[])arg).Length];
                Array.Copy((sbyte[])arg, buff, buff.Length);
                stream.WriteBytes(buff);
            }
            else if (type == typeof(bool))
            {
                stream.WriteBool((bool)arg);
            }
            else if (typeof(ProtoBuf.IExtensible).IsAssignableFrom(type))
            {
                //stream.WriteString(arg.GetType().FullName);

                // +2 是由于WriteBuf存在Length的长度
                var segment = new ArraySegment<byte>(stream.Buf, stream.WPos + 2, stream.Capicity - stream.WPos - 2);

                var pstream = new System.IO.MemoryStream(segment.Array, segment.Offset, segment.Count);
                ProtoBuf.Meta.RuntimeTypeModel.Default.Serialize(pstream, arg);
                stream.WriteUInt16((ushort)pstream.Position);
                stream.WPos += (int)pstream.Position;

                //var buf = new byte[stream.WPos];
                //Buffer.BlockCopy(stream.Buf, 0, buf, 0, stream.WPos);
                //stream.WriteBytes(buf);
            }
            else if (typeof(IDictionary).IsAssignableFrom(type))
            {
                var dict = (arg as IDictionary);
                stream.WriteUInt16((ushort)dict.Count);
                var enumer = dict.GetEnumerator();
                while (enumer.MoveNext())
                {
                    var item = enumer.Current;
                    var itemKey = item.GetType().GetProperty("Key").GetValue(item, null);
                    var itemValue = item.GetType().GetProperty("Value").GetValue(item, null);

                    doPack(stream, itemKey.GetType(), itemKey);
                    doPack(stream, itemValue.GetType(), itemValue);
                }
                //var buf = new byte[stream.WPos];
                //Buffer.BlockCopy(stream.Buf, 0, buf, 0, stream.WPos);
                //stream.WriteBytes(buf);
            }
            else if (typeof(IList).IsAssignableFrom(type))
            {
                var list = (arg as IList);
                stream.WriteUInt16((ushort)list.Count);
                for (var i = 0; i < list.Count; i++)
                {
                    var listitem = list[i];
                    doPack(stream, listitem.GetType(), listitem);
                }

                //var buf = new byte[stream.WPos];
                //Buffer.BlockCopy(newstream.Buf, 0, buf, 0, newstream.WPos);
                //stream.WriteBytes(buf);
            }
            else if (type.IsArray)
            {
                var fields = arg.GetType().GetFields();

                stream.WriteUInt16((ushort)fields.Length);

                for (var i = 0; i < fields.Length; i++)
                {
                    var field = fields[i];
                    //var index = arg.GetType().IsArray ? new object[] { i } : null;
                    doPack(stream, field.FieldType, field.GetValue(arg));
                }
            }
            else if (type.IsEnum)
            {
                var fields = arg.GetType().GetFields();

                var field = fields[0];
                //var index = arg.GetType().IsArray ? new object[] { i } : null;
                doPack(stream, field.FieldType, field.GetValue(arg));
            }
            else if (type.IsValueType || type.IsClass)
            {
                //if (arg.GetType().IsArray)
                //{
                //    //array
                //}
                //else
                //{
                //    //struct
                //    var props = arg.GetType().GetProperties();

                //    for (var i = 0; i < props.Length; i++)
                //    {
                //        var prop = props[i];
                //        stream.WriteBytes(doPack(prop.GetValue(arg, null)));
                //    }
                //}
                var fields = arg.GetType().GetFields();

                for (var i = 0; i < fields.Length; i++)
                {
                    var field = fields[i];
                    //var index = arg.GetType().IsArray ? new object[] { i } : null;
                    doPack(stream, field.FieldType, field.GetValue(arg));
                }

                //var buf = new byte[newstream.WPos];
                //Buffer.BlockCopy(newstream.Buf, 0, buf, 0, newstream.WPos);
                //stream.WriteBytes(buf);
            }
            else
            {
                throw new NotSupportedException("Unknow rpc method param type: [" + arg.GetType().Name + "]");
            }
        }

        public static bool UnpackMethod(MethodInfo minfo, byte[] buf, out object[] args)
        {
            var argList = new List<object>();
            var stream = MMStreamPool.GetTmp();// new MMStream(buf);
            stream.Reset(buf);
            var paramarr = minfo.GetParameters();
            for(var i = 0; i < paramarr.Length; i++)
            {
                var paraminfo = paramarr[i];
                argList.Add(doUnpack(stream, paraminfo.ParameterType));
            }
            args = argList.ToArray();
            MMStreamPool.PutTmp(stream);
            return true;
        }

        public static object UnpackParam(Type type, byte[] buf)
        {
            var stream = MMStreamPool.GetTmp();//new MMStream(buf);
            stream.Reset(buf);
            var ret = doUnpack(stream, type);
            MMStreamPool.PutTmp(stream);
            return ret;
        }

        private static object doUnpack(MMStream stream, Type type)
        {
            //if (stream.IsEOF())
            //    return null;

            //Debug.Log("doUnpack type:" + type + " type.IsValueType:" + type.IsValueType + " type.IsArray:" + type.IsArray);

            if (!stream.IsEOF())
            {
                if (type == typeof(byte))
                {
                    return stream.ReadByte();
                }
                else if (type == typeof(sbyte))
                {
                    return (sbyte)stream.ReadByte();
                }
                else if (type == typeof(ushort))
                {
                    return stream.ReadUInt16();
                }
                else if (type == typeof(uint))
                {
                    return stream.ReadUInt32();
                }
                else if (type == typeof(ulong))
                {
                    return stream.ReadUInt64();
                }
                else if (type == typeof(short))
                {
                    return stream.ReadInt16();
                }
                else if (type == typeof(int))
                {
                    return stream.ReadInt32();
                }
                else if (type == typeof(long))
                {
                    return stream.ReadInt64();
                }
                else if (type == typeof(float))
                {
                    return stream.ReadFloat();
                }
                else if (type == typeof(double))
                {
                    return stream.ReadDouble();
                }
                else if (type == typeof(string))
                {
                    return stream.ReadString();
                }
                else if (type == typeof(byte[]))
                {
                    return stream.ReadBytes();
                }
                else if (type == typeof(sbyte[]))
                {
                    var buff = stream.ReadBytes();
                    var sbuff = new sbyte[buff.Length];
                    Array.Copy(buff, sbuff, buff.Length);
                    return sbuff;
                }
                else if (type == typeof(bool))
                {
                    return stream.ReadBoolean();
                }
                else if (type.IsArray)
                {
                    var count = stream.ReadUInt16();
                    object newobj = Array.CreateInstance(type.GetElementType(), count);

                    for (var i = 0; i < count; i++)
                    {
                        var val = doUnpack(stream, type.GetElementType());
                        newobj.GetType().GetMethod("SetValue", new Type[2] { typeof(object), typeof(int) }).Invoke(newobj, new object[] { val, i });
                    }

                    return newobj;
                }
                else if (typeof(IDictionary).IsAssignableFrom(type))
                {
                    var count = stream.ReadUInt16();
                    var eletypes = type.GetGenericArguments();
                    var newdict = Activator.CreateInstance(type) as IDictionary;

                    for (var i = 0; i < count; i++)
                    {
                        var key = doUnpack(stream, eletypes[0]);
                        var val = doUnpack(stream, eletypes[1]);
                        newdict[key] = val;
                    }
                    return newdict;
                }
                else if (typeof(IList).IsAssignableFrom(type))
                {
                    var count = stream.ReadUInt16();
                    var newlist = Activator.CreateInstance(type) as IList;
                    var eletype = type.GetGenericArguments()[0];//.GetElementType();

                    for (var i = 0; i < count; i++)
                    {
                        newlist.Add(doUnpack(stream, eletype));
                    }
                    return newlist;
                }
            }
            
            if (typeof(ProtoBuf.IExtensible).IsAssignableFrom(type))
            {
                System.IO.MemoryStream pstream = null;
                if (stream.IsEOF())
                {
                    pstream = new System.IO.MemoryStream();
                }
                else
                {
                    var msgBuf = stream.ReadBytes();

                    pstream = new System.IO.MemoryStream(msgBuf, 0, msgBuf.Length);
                    pstream.SetLength(msgBuf.Length);
                }
                    
                return ProtoBuf.Meta.RuntimeTypeModel.Default.Deserialize(pstream, null, type);
            }
            else if (type.IsEnum)
            {
                object newobj = Activator.CreateInstance(type);

                if (!stream.IsEOF())
                {
                    var fields = newobj.GetType().GetFields();

                    var field = fields[0];
                    var val = doUnpack(stream, field.FieldType);
                    //var index = type.IsArray ? new object[] { i } : null;
                    field.SetValue(newobj, val);
                }
                    
                return newobj;
            }
            else if (type.IsValueType || type.IsClass)
            {
                object newobj = Activator.CreateInstance(type);

                if (!stream.IsEOF())
                {
                    var fields = newobj.GetType().GetFields();

                    for (var i = 0; i < fields.Length; i++)
                    {
                        var field = fields[i];
                        var val = doUnpack(stream, field.FieldType);
                        //var index = type.IsArray ? new object[] { i } : null;
                        field.SetValue(newobj, val);
                    }
                }
                    
                return newobj;
            }
            else
            {
                throw new NotSupportedException("Unknow rpc method param type: [" + type.Name + "]");
            }
        }

        /// <summary>
        /// 参数打包
        /// </summary>
        /// <param name="args">参数名</param>
        /// <param name="msgBuf">需要返回的流</param>
        /// <returns>是否打包成功</returns>
        public static bool ArgsPack(object[] args, out byte[] msgBuf)
        {
            var stream = MMStreamPool.Get();
            stream.WPos = 0;
            stream.RPos = 0;

            try
            {
                foreach (var arg in args)
                {
                    if (arg is byte)
                    {
                        stream.WriteByte((byte)ArgType.typeUint8);
                        stream.WriteByte((byte)arg);
                    }
                    else if (arg is ushort)
                    {
                        stream.WriteByte((byte)ArgType.typeUint16);
                        stream.WriteUInt16((ushort)arg);
                    }
                    else if (arg is uint)
                    {
                        stream.WriteByte((byte)ArgType.typeUint32);
                        stream.WriteUInt32((uint)arg);
                    }
                    else if (arg is ulong)
                    {
                        stream.WriteByte((byte)ArgType.typeUint64);
                        stream.WriteUInt64((ulong)arg);
                    }
                    else if (arg is short)
                    {
                        stream.WriteByte((byte)ArgType.typeInt16);
                        stream.WriteInt16((short)arg);
                    }
                    else if (arg is int)
                    {
                        stream.WriteByte((byte)ArgType.typeInt32);
                        stream.WriteInt32((int)arg);
                    }
                    else if (arg is long)
                    {
                        stream.WriteByte((byte)ArgType.typeInt64);
                        stream.WriteInt64((long)arg);
                    }
                    else if (arg is float)
                    {
                        stream.WriteByte((byte)ArgType.typeFloat32);
                        stream.WriteFloat((float)arg);
                    }
                    else if (arg is double)
                    {
                        stream.WriteByte((byte)ArgType.typeFloat64);
                        stream.WriteDouble((double)arg);
                    }
                    else if (arg is string)
                    {
                        stream.WriteByte((byte)ArgType.typeString);
                        stream.WriteString((string)arg);
                    }
                    else if (arg.GetType() == typeof(byte[]))
                    {
                        stream.WriteByte((byte)ArgType.typeBytes);
                        stream.WriteBytes((byte[])arg);
                    }
                    else if (arg is bool)
                    {
                        stream.WriteByte((byte)ArgType.typeBool);
                        stream.WriteBool((bool)arg);
                    }
                    else
                    {
                        //throw new NotSupportedException(
                        //    "Unknow rpc method param type: [" + arg.GetType().Name + "]");
                        //var msgInfo = MsgService.Instance.GetMsgByName(arg.GetType().Name);
                        //if (msgInfo == null)
                        //{
                        //    throw new NotSupportedException(
                        //        "Unknow msg info , msg type is [" + arg.GetType().Name + "]");
                        //}

                        stream.WriteByte((byte)ArgType.typeProto);
                        //stream.WriteUInt16((ushort)msgInfo.Id);
                        stream.WriteString(arg.GetType().FullName);

                        // +2 是由于WriteBuf存在Length的长度
                        var segment = new ArraySegment<byte>(stream.Buf, stream.WPos + 2, stream.Capicity - stream.WPos - 2);
                        //PacketPolicy.Packed(ref segment, msgInfo, arg);
                        //stream.WriteUInt16((ushort)segment.Count);
                        //stream.WPos += segment.Count;

                        var pstream = new System.IO.MemoryStream(segment.Array, segment.Offset, segment.Count);
                        ProtoBuf.Meta.RuntimeTypeModel.Default.Serialize(pstream, arg);
                        stream.WriteUInt16((ushort)segment.Count);
                        stream.WPos += segment.Count;
                    }
                }
            }
            catch (Exception)
            {
                msgBuf = null;
                return false;
            }

            if (stream.WPos > 0)
            {
                msgBuf = new byte[stream.WPos];
                Buffer.BlockCopy(stream.Buf, 0, msgBuf, 0, stream.WPos);
            }
            else
            {
                msgBuf = null;
            }
            MMStreamPool.Put(stream);

            return true;
        }
                
        /// <summary>
        /// 参数解包
        /// </summary>
        /// <param name="buf">需要解包的流</param>
        /// <param name="args">解包后的参数表</param>
        /// <param name="protoBinaryReserved">是否进行proto二进制解决</param>
        /// <returns>是否成功解包</returns>
        public static bool ArgsUnpack(byte[] buf, out object[] args, bool protoBinaryReserved = false)
        {
            if (buf == null)
            {
                args = new object[] { };
                return true;
            }

            var unPackSucceed = true;
            var stream = MMStreamPool.GetTmp(); //new MMStream(buf);
            stream.Reset(buf);
            var argList = new List<object>();

            while (!stream.IsEOF())
            {
                var type = (ArgType)stream.ReadByte();
                switch (type)
                {
                    case ArgType.typeInt8:
                    case ArgType.typeUint8:
                        argList.Add(stream.ReadByte());
                        break;
                    case ArgType.typeInt16:
                        argList.Add(stream.ReadInt16());
                        break;
                    case ArgType.typeInt32:
                        argList.Add(stream.ReadInt32());
                        break;
                    case ArgType.typeInt64:
                        argList.Add(stream.ReadInt64());
                        break;
                    case ArgType.typeUint16:
                        argList.Add(stream.ReadUInt16());
                        break;
                    case ArgType.typeUint32:
                        argList.Add(stream.ReadUInt32());
                        break;
                    case ArgType.typeUint64:
                        argList.Add(stream.ReadUInt64());
                        break;
                    case ArgType.typeBool:
                        argList.Add(stream.ReadBoolean());
                        break;
                    case ArgType.typeString:
                        argList.Add(stream.ReadString());
                        break;
                    case ArgType.typeBytes:
                        argList.Add(stream.ReadBytes());
                        break;
                    case ArgType.typeFloat32:
                        argList.Add(stream.ReadFloat());
                        break;
                    case ArgType.typeFloat64:
                        argList.Add(stream.ReadDouble());
                        break;
                    case ArgType.typeProto:
                        //throw new NotSupportedException(
                        //    "not support proto param for rpc method");
                        object msg;
                        unPackSucceed = UnpackProto(stream, protoBinaryReserved, out msg);
                        argList.Add(msg);
                        break;
                }
            }

            args = argList.ToArray();
            MMStreamPool.PutTmp(stream);
            return unPackSucceed;
        }

        /// <summary>
        /// 解包Proto
        /// </summary>
        /// <param name="stream">流</param>
        /// <param name="protoBinaryReserved">是否以二进制形式解决</param>
        /// <param name="msg">消息包</param>
        /// <returns>是否根据规定方案解包</returns>
        private static bool UnpackProto(MMStream stream, bool protoBinaryReserved, out object msg)
        {
            var fullname = (string)stream.ReadString();
            var msgBuf = stream.ReadBytes();

            //var msgId = (int)stream.ReadUInt16();
            
            //var msgInfo = MsgService.Instance.GetMsgById(msgId);

            //if (protoBinaryReserved)
            //{
            //    //if (msgInfo == null)
            //    //{
            //    //    msg = msgBuf;
            //    //    return false;
            //    //}

            //    msg = new RawProto { name = msgInfo.Name, rawData = msgBuf };
            //    return true;
            //}

            var ProtocolType = Type.GetType(fullname);
            var pstream = new System.IO.MemoryStream(msgBuf, 0, msgBuf.Length);
            pstream.SetLength(msgBuf.Length);
            msg = ProtoBuf.Meta.RuntimeTypeModel.Default.Deserialize(pstream, null, ProtocolType);

            //msg = PacketPolicy.Unpacked(msgInfo, new ArraySegment<byte>(msgBuf, 0, msgBuf.Length));
            return true;
        }
    }
}
