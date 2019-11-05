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
using System.Threading;
using GameBox.Network;

namespace GameBox.Service
{
    /// <summary>
    /// 盖亚网络系统 打包解包器
    /// </summary>
    internal sealed class GaiaPacker : IPacker
    {
        /// <summary>
        /// 发送缓冲区
        /// </summary>
        private readonly byte[] sentBuffer;

        private readonly byte[] compSentBuffer;

        /// <summary>
        /// 打包处理策略
        /// </summary>
        private readonly IPackerPolicy packerPolicy;

        /// <summary>
        /// RpcId自增器
        /// </summary>
        private static int rpcIdIncreased;

        private IMsgDefined msgDefined;

        /// <summary>
        /// 构建一个盖亚网络系统 打包解包器实例
        /// </summary>
        public GaiaPacker(IPackerPolicy packerPolicy, IMsgDefined msgdefined)
        {
            sentBuffer = new byte[Header.MaxPacketSize];
            compSentBuffer = new byte[Header.MaxPacketSize];
            this.packerPolicy = packerPolicy;
            this.msgDefined = msgdefined;
        }

        /// <summary>
        /// 打包
        /// </summary>
        /// <param name="packet">输入的协议对象</param>
        /// <returns>需要发送的字节流</returns>
        public ArraySegment<byte> Pack(object packet)
        {
            //            offset             
            // ---------------------------------------------------------
            // |    包头    |              包体（Count）               |
            // ---------------------------------------------------------
            // |              ArraySegment(包头+包体)                  |
            // ---------------------------------------------------------
            // 返回的ArraySegment，Offset为包头偏移位置，Count为包体长度

            var msgInfo = msgDefined.GetMsgByName(packet.GetType().FullName);
            if (msgInfo == null)
            {
                throw new NotSupportedException("Can not find msg info , msg name is [" + packet.GetType().FullName + "]");
            }

            var sent = new ArraySegment<byte>(sentBuffer, Header.HeadSize, sentBuffer.Length - Header.HeadSize);

            packerPolicy.Packed(ref sent, msgInfo, packet);

            var header = new Header
            {
                MsgId = msgInfo.Id,
            };

            header.FillMsgId(sent);
            var datalen = sent.Count;
            var destbuff = sentBuffer;

            header.Encrypt = true;
            if (datalen >= 100)
            {
                //compress
                header.Compress = true;
                destbuff = compSentBuffer;
                datalen = Snappy.Compress(sent.Array, sent.Offset - 2, sent.Count + 2, compSentBuffer, Header.HeadSize - 2);
                //Debug.LogError("send data compressed pre count->" + sent.Count + " after->" + datalen);
                header.BodyLength = datalen;
            }
            else
            {
                header.BodyLength = datalen + 2;
            }

            sent = new ArraySegment<byte>(destbuff, 0, header.BodyLength + Header.HeadSize - 2);
            header.Fill(sent);
            XorEncrypt.Encrypt(destbuff, Header.HeadSize - 2, header.BodyLength);

            return sent;
        }
        public string ModuleName;
        public object Unpack(ArraySegment<byte> packet)
        {
            var header = new Header(packet);
            
            if (header.Encrypt)
            {
                XorEncrypt.Decrypt(packet.Array, Header.HeadSize - 2, header.BodyLength);
            }

            if (header.Compress)
            {
                var uncompdata = Snappy.Uncompress(packet.Array, Header.HeadSize - 2, header.BodyLength);
                header.Body = new ArraySegment<byte>(uncompdata, 2, uncompdata.Length - 2);
                header.MsgId = uncompdata[0] | (uint)uncompdata[1] << 8;
            }
            else if (header.Encrypt)
            {
                header.MsgId = packet.Array[packet.Offset + 4] | (uint)packet.Array[packet.Offset + 5] << 8;
            }

            //if (ModuleName == "GameBox.Cratos.GBoxSync")
            //    Debug.LogError("Unpack msgid->" + header.MsgId + " buff len->" + packet.Count + " header.BodyLength->" + header.BodyLength + " header.Encrypt->" + header.Encrypt + " header.Compress->" + header.Compress);

            var msgInfo = msgDefined.GetMsgById(header.MsgId);
            if (msgInfo == null)
            {
                throw new NotSupportedException("Can not find msg info , msg id is [" + header.MsgId + "]");
            }

            var result = packerPolicy.Unpacked(msgInfo, header.Body);
            if (msgInfo.Id != MsgConst.SyncMsgID)
            {
                return result;
            }

            //MsgService.SyncMsg(result);
            return null;
        }
    }
}
