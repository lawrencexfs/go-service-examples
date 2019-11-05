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

namespace GameBox.Service
{
    /// <summary>
    /// 包头数据
    /// </summary>
    internal struct Header
    {
        /// <summary>
        /// 包头长度(字节)
        /// </summary>
        public static int HeadSize
        {
            get
            {
                return 6;
            }
        }

        /// <summary>
        /// 最大整包长度(包头+包体)(字节)
        /// </summary>
        public static int MaxPacketSize
        {
            get { return 65536; }
        }

        /// <summary>
        /// 消息Id
        /// </summary>
        public uint MsgId { get; set; }

        /// <summary>
        /// 是否被压缩
        /// </summary>
        public bool Compress { get; set; }

        /// <summary>
        /// 是否处于加密
        /// </summary>
        public bool Encrypt { get; set; }

        /// <summary>
        /// 包体长度
        /// </summary>
        public int BodyLength { get; set; }

        /// <summary>
        /// 包体数据
        /// </summary>
        public ArraySegment<byte> Body { get; set; }

        /// <summary>
        /// 初始化一个包头
        /// </summary>
        /// <param name="data">数据</param>
        public Header(ArraySegment<byte> data) : this()
        {
            BodyLength = (data.Array[data.Offset + 0] | data.Array[data.Offset + 1] << 8 | data.Array[data.Offset + 2] << 16);
            Compress = (data.Array[data.Offset + 3] & (1 << 0)) > 0;
            Encrypt = (data.Array[data.Offset + 3] & (1 << 1)) > 0;
            MsgId = data.Array[data.Offset + 4] | (uint)data.Array[data.Offset + 5] << 8;

            Body = new ArraySegment<byte>(data.Array, data.Offset + HeadSize, data.Count - HeadSize);
        }

        /// <summary>
        /// 填充数据
        /// </summary>
        /// <param name="data">字节流</param>
        public void Fill(ArraySegment<byte> data)
        {
            data.Array[data.Offset + 0] = (byte)BodyLength;
            data.Array[data.Offset + 1] = (byte)(BodyLength >> 8);
            data.Array[data.Offset + 2] = (byte)(BodyLength >> 16);

            data.Array[data.Offset + 3] = 0;

            if (Compress)
            {
                data.Array[data.Offset + 3] |= 0x1;
            }

            if (Encrypt)
            {
                data.Array[data.Offset + 3] |= 0x2;
            }
        }

        public void FillMsgId(ArraySegment<byte> data)
        {
            data.Array[4] = (byte)MsgId;
            data.Array[5] = (byte)(MsgId >> 8);
        }
    }
}
