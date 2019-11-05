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
using System.ComponentModel.Design.Serialization;
using System.Linq;
using System.Text;
using System.Threading;

namespace GameBox.Service
{
    internal class MsgGenerate<T> where T : IMsg, new()
    {
        public static IMsg Generate()
        {
            return new T();
        }

        public void Recycle()
        {

        }
    }

    //由于底层修改了网络逻辑，改成了多线程，所以静态方法无法使用了。但是效果有待观察，所以只修改父类
    internal class MsgGenerateStatic<T> where T : IMsg, new()
    {
        //public static T instance;

        public static IMsg Generate()
        {
            //if (instance == null)
            //{
            //instance = GObjectPool.Create<T>();
            //}
            //return instance;

            return GObjectPool.Create<T>();
        }

        public void Recycle()
        {
            GObjectPool.Recycle(this);
        }

    }



    internal class ClientVertifyReq : MsgGenerate<ClientVertifyReq>, IMsg
    {
        public byte Source = 0;
        public UInt64 UID;
        public byte[] Token = new byte[32];

        public void Marshal(MMStream stream)
        {
            stream.WriteByte(Source);
            stream.WriteUInt64(UID);
            stream.WriteBytes(Token);
        }

        public void UnMarshal(MMStream stream)
        {
            Source = stream.ReadByte();
            UID = stream.ReadUInt64();
            Token = stream.ReadBytes(Token.Length);
        }
    }

    internal class UserDuplicateLoginNotify : MsgGenerate<UserDuplicateLoginNotify>, IMsg
    {

        public void Marshal(MMStream stream)
        {
        }

        public void UnMarshal(MMStream stream)
        {
        }
    }


    internal class ClientVertifySucceedRet : MsgGenerate<ClientVertifySucceedRet>, IMsg
    {
        public byte Source;
        public UInt64 UID;

        public UInt64 SourceID;
        public byte Type;


        public void Marshal(MMStream stream)
        {
            stream.WriteByte(Source);
            stream.WriteUInt64(UID);
            stream.WriteUInt64(SourceID);
            stream.WriteByte(Type);
        }

        public void UnMarshal(MMStream stream)
        {
            Source = stream.ReadByte();
            UID = stream.ReadUInt64();
            SourceID = stream.ReadUInt64();
            Type = stream.ReadByte();
        }

    }

    internal class ClientVertifyFailedRet : MsgGenerate<ClientVertifyFailedRet>, IMsg
    {
        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {

        }
    }

    //internal class Ping : MsgGenerateStatic<Ping>, IMsg
    //{
    //    public uint Id;
    //    public uint Timestamp;

    //    public void Marshal(MMStream stream)
    //    {
    //        stream.WriteUInt32(Id);
    //        stream.WriteUInt32(Timestamp);
    //    }

    //    public void UnMarshal(MMStream stream)
    //    {
    //        Id = stream.ReadUInt32();
    //        Timestamp = stream.ReadUInt32();
    //    }
    //}

    internal class HeartBeat : MsgGenerateStatic<HeartBeat>, IMsg
    {
        public void Marshal(MMStream stream)
        {
        }

        public void UnMarshal(MMStream stream)
        {

        }
    }

    internal class HeartBeatResponse : MsgGenerateStatic<HeartBeatResponse>, IMsg
    {
        public void Marshal(MMStream stream)
        {
        }

        public void UnMarshal(MMStream stream)
        {
        }
    }

    internal class MsgSyncRet : MsgGenerate<MsgSyncRet>, IMsg
    {
        public string msgSync;

        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            //msgSync = Encoding.UTF8.GetString(stream.GetUsedBytes());
            msgSync = Encoding.UTF8.GetString(stream.Buf, stream.RPos, stream.Capicity - stream.RPos);
        }
    }

    internal class RPCMsg : MsgGenerate<RPCMsg>, IMsg
    {
        public byte srvType;
        public UInt64 entityID;
        public string methodName;
        public byte[] args;

        public void Marshal(MMStream stream)
        {
            stream.WriteByte(srvType);
            stream.WriteUInt64(entityID);
            stream.WriteString(methodName);
            stream.WriteBytes(args);

        }

        public void UnMarshal(MMStream stream)
        {
            srvType = stream.ReadByte();
            entityID = stream.ReadUInt64();
            methodName = stream.ReadString();
            args = stream.ReadBytes();
        }
    }


    internal class EnterSpace : MsgGenerateStatic<EnterSpace>, IMsg
    {
        public UInt64 spaceID;
        public string mapName;
        public UInt64 entityID;
        public string outAddr;
        public UInt32 timeStamp;

        public void Marshal(MMStream stream)
        {
            stream.WriteUInt64(spaceID);
            stream.WriteString(mapName);
            stream.WriteUInt64(entityID);
            stream.WriteString(outAddr);
            stream.WriteUInt32(timeStamp);
        }

        public void UnMarshal(MMStream stream)
        {
            spaceID = stream.ReadUInt64();
            mapName = stream.ReadString();
            entityID = stream.ReadUInt64();
            outAddr = stream.ReadString();
            timeStamp = stream.ReadUInt32();
        }
    }


    internal class LeaveSpace : MsgGenerateStatic<LeaveSpace>, IMsg
    {
        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {

        }
    }


    internal class EnterAOI : MsgGenerateStatic<EnterAOI>, IMsg
    {
        public UInt64 entityID;
        public string entityType;
        public byte[] state;
        public UInt16 propNum;
        public byte[] properties;
        public EntityBaseProps baseProps;

        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            entityID = stream.ReadUInt64();
            entityType = stream.ReadString();
            state = stream.ReadBytes();
            propNum = stream.ReadUInt16();
            properties = stream.ReadBytes();

            var len = stream.ReadUInt16();
            if (len == 0)
            {
                baseProps = null;
            }
            else
            {
                baseProps = new EntityBaseProps();
                baseProps.UnMarshal(stream);

            }
        }
    }

    internal class LeaveAOI : MsgGenerateStatic<LeaveAOI>, IMsg
    {
        public UInt64 entityID;

        public void Marshal(MMStream stream)
        {
            stream.WriteUInt64(entityID);

        }

        public void UnMarshal(MMStream stream)
        {
            entityID = stream.ReadUInt64();

        }
    }

    internal class AOIPosChange : MsgGenerateStatic<AOIPosChange>, IMsg
    {

        public int num = 0;
        public byte[] data;


        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            num = stream.ReadInt32();
            data = stream.ReadBytes();


        }

        public int getNum()
        {
            return num;
        }
    }


    internal class SpaceEntityMsg : MsgGenerateStatic<SpaceEntityMsg>, IMsg
    {
        public byte[] data;

        public void Marshal(MMStream stream)
        {
            stream.WriteBytes(data);

        }

        public void UnMarshal(MMStream stream)
        {

        }
    }

    internal class PropSyncClient : MsgGenerateStatic<PropSyncClient>, IMsg
    {
        public UInt64 entityID;
        public UInt32 num;
        public FixedLengthArray<byte> data = new FixedLengthArray<byte>(1024);


        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            entityID = stream.ReadUInt64();
            num = stream.ReadUInt32();
            //data = stream.ReadBytes();

            if (num > 1024)
            {
                throw new ArgumentNullException("PropSyncClient UnMarshal");
            }

            stream.ReadFixByteArray(ref data);
        }
    }

    internal class MainEntityGenerate : MsgGenerate<MainEntityGenerate>, IMsg
    {
        public UInt64 entityID;
        public UInt32 propNum;
        public byte[] properties;

        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            entityID = stream.ReadUInt64();
            propNum = stream.ReadUInt32();
            properties = stream.ReadBytes();

        }
    }


    internal class SpaceUserConnect : MsgGenerate<SpaceUserConnect>, IMsg
    {
        public UInt64 UID;
        public UInt64 SpaceID;


        public void Marshal(MMStream stream)
        {
            stream.WriteUInt64(UID);
            stream.WriteUInt64(SpaceID);

        }

        public void UnMarshal(MMStream stream)
        {


        }
    }

    internal class SpaceUserConnectSucceedRet : MsgGenerate<SpaceUserConnectSucceedRet>, IMsg
    {


        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {

        }
    }

    internal class SyncClock : MsgGenerate<SyncClock>, IMsg
    {
        public UInt32 timeStamp;


        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            timeStamp = stream.ReadUInt32();

        }
    }

    internal class SyncUserState : MsgGenerateStatic<SyncUserState>, IMsg
    {
        public ulong id;
        //public byte[] data = new byte[1024];
        public FixedLengthArray<byte> statdata = new FixedLengthArray<byte>(1024);

        public void append(byte[] data, int count)
        {
            //this.data = new byte[count];
            //Array.Copy(data, data, count);
            //data.Copy(data,count);
            statdata.Clear();
            statdata.Set(data, count);
        }

        public void Marshal(MMStream stream)
        {
            stream.WriteUInt64(id);
            stream.WriteBuf(statdata.getArray(), statdata.usedLength);
        }

        public void UnMarshal(MMStream stream)
        {
            id = stream.ReadUInt64();

            int len = stream.ReadUInt16();
            var marbytearr = statdata.getArray();
            stream.ReadBytes(ref marbytearr, len);

            //data = stream.ReadBytes();
        }
    }

    internal class AOISyncUserState : MsgGenerateStatic<AOISyncUserState>, IMsg
    {
        public uint num;

        public FixedLengthArray<UInt64> eids = new FixedLengthArray<UInt64>(128);
        public FixedLengthArray<FixedLengthArray<byte>> eds = new FixedLengthArray<FixedLengthArray<byte>>(128);

        //public List<UInt64> eids = new List<ulong>();
        //public List<byte[]> eds = new List<byte[]>();

        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            num = stream.ReadUInt32();

            eids.Clear();

            eds.SetUsedLength((int)num);

            for (int i = 0; i < num; i++)
            {
                eids.Add(stream.ReadUInt64());

                //int blen = stream.ReadUInt16();

                var edsbytearr = eds.getArray()[i];
                if (edsbytearr == null)
                {
                    eds.getArray()[i] = new FixedLengthArray<byte>(2048);
                    edsbytearr = eds.getArray()[i];
                }

                stream.ReadFixByteArray(ref edsbytearr);


            }

        }
    }


    internal class AdjustUserState : MsgGenerateStatic<AdjustUserState>, IMsg
    {
        public UInt64 entityId;
        public UInt32 adjustCount;

        public FixedLengthArray<byte> stateData = new FixedLengthArray<byte>(512);

        public void Marshal(MMStream stream)
        {
        }

        public void UnMarshal(MMStream stream)
        {
            entityId = stream.ReadUInt64();
            adjustCount = stream.ReadUInt32();
            stream.ReadFixByteArray(ref stateData);



            //data = stream.ReadBytes();
        }
    }


    internal class AdjustStateRsp : MsgGenerateStatic<AdjustStateRsp>, IMsg
    {
        public UInt64 entityID;
        public UInt32 adjustCount;

        public void Marshal(MMStream stream)
        {
            stream.WriteUInt64(entityID);
            stream.WriteUInt32(adjustCount);
        }

        public void UnMarshal(MMStream stream)
        {

        }
    }


    //据说，这个消息后面会被干掉。只要留 ENTERAOI和LEAVEAOI消息
    internal class EntityAOIS : MsgGenerateStatic<EntityAOIS>, IMsg
    {
        public int num;

        //public List<byte[]> data = new List<byte[]>();
        //缓冲区大小 512*512 bytes
        public FixedLengthArray<FixedLengthArray<byte>> aoi_data = new FixedLengthArray<FixedLengthArray<byte>>(128);

        public EntityAOIS()
        {
            for (int i = 0; i < 128; i++)
            {
                aoi_data.Add(new FixedLengthArray<byte>(2048));
            }
        }
        public void Marshal(MMStream stream)
        {

        }

        public void UnMarshal(MMStream stream)
        {
            num = stream.ReadInt32();

            if (num < aoi_data.totalLength)
            {
                aoi_data.SetUsedLength(num);
                for (int i = 0; i < num; i++)
                {
                    var aoiarr = aoi_data.getArray()[i];
                    stream.ReadFixByteArray(ref aoiarr);
                }
            }
        }
    }

    internal class EntityBaseProps : MsgGenerate<EntityBaseProps>, IMsg
    {
        static readonly int maxLinkersLength = 8;
        static readonly int maxCommersLength = 8;

        /// <summary>
        /// EntityID
        /// </summary>
        public ulong id;

        /// <summary>
        /// 链接到哪个Entity
        /// </summary>
        public ulong linkTarget;

        /// <summary>
        /// 哪些Entity链接到我
        /// </summary>
        public FixedLengthArray<ulong> linkers = new FixedLengthArray<ulong>(maxLinkersLength);

        /// <summary>
        /// 委托给哪个Entity
        /// </summary>
        public ulong commTarget;

        /// <summary>
        /// 哪些Entity委托给我
        /// </summary>
        public FixedLengthArray<ulong> commers = new FixedLengthArray<ulong>(maxCommersLength);

        public void Marshal(MMStream stream) { }

        public void UnMarshal(MMStream stream)
        {
            id = stream.ReadUInt64();

            linkTarget = stream.ReadUInt64();

            var linkLength = stream.ReadByte();
            if (linkLength <= linkers.totalLength)
            {
                linkers.Clear();
            }
            else
            {
                linkers = new FixedLengthArray<ulong>(linkLength);
            }
            for (int i = 0; i < linkLength; i++)
            {
                linkers.Add(stream.ReadUInt64());
            }

            commTarget = stream.ReadUInt64();

            var commLength = stream.ReadByte();
            if (commLength <= commers.totalLength)
            {
                commers.Clear();
            }
            else
            {
                commers = new FixedLengthArray<ulong>(commLength);
            }
            for (int i = 0; i < commLength; i++)
            {
                commers.Add(stream.ReadUInt64());
            }
        }
    }

    /// <summary>
    /// 实体事件
    /// </summary>
    internal class EntityEvent : MsgGenerate<EntityEvent>, IMsg
    {
        public UInt64 entityID;
        public string methodName;
        public byte[] args;


        public void Marshal(MMStream stream)
        {
            stream.WriteUInt64(entityID);
            stream.WriteString(methodName);
            stream.WriteBytes(args);

        }

        public void UnMarshal(MMStream stream)
        {
            entityID = stream.ReadUInt64();
            methodName = stream.ReadString();
            args = stream.ReadBytes();
        }
    }
}
