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
using GameBox.Channel;

namespace GameBox.Service
{
    /// <summary>
    /// 盖亚网络系统 协议分包逻辑
    /// </summary>
    internal sealed class GaiaFragment : IFragment
    {
        /// <summary>
        /// 最大单个数据包的尺寸
        /// </summary>
        public int MaxPackageSize
        {
            get
            {
                return Header.MaxPacketSize;
            }
        }

        /// <summary>
        /// 超过这个大小会进行数据压缩（字节）
        /// </summary>
        private int MinCompressSize
        {
            get
            {
                return 100 + Header.HeadSize;
            }
        }

        /// <summary>
        /// 是否使用加密
        /// </summary>
        private bool UseEncrypt
        {
            get
            {
                return false;
            }
        }

        /// <summary>
        /// 压缩缓冲区
        /// </summary>
        private readonly byte[] receiveCompressBuffer;

        /// <summary>
        /// 发送缓冲区
        /// </summary>
        private readonly byte[] sendCompressBuffer;

        private const int HeadSize = 6;

        /// <summary>
        /// 构造一个新的协议分包逻辑实例
        /// </summary>
        public GaiaFragment()
        {
            //receiveCompressBuffer = new byte[MaxPackageSize];
            //sendCompressBuffer = new byte[MaxPackageSize];
        }

        /// <summary>
        /// <para>检查包的完整性</para>
        /// <para>如果能够得到包长，则返回包的在buffer中的长度(包含包头)，否则返回0继续等待数据</para>
        /// <para>如果协议有问题，则填入ex参数，当前连接会因此断开</para>
        /// </summary>
        /// <param name="source">输入数据</param>
        /// <param name="ex">协议异常</param>
        /// <returns>包长度</returns>
        public int Input(ArraySegment<byte> source, out Exception ex)
        {
            ex = null;
            try
            {
                if (source.Count < HeadSize)
                {
                    return 0;
                }
                var len = (source.Array[source.Offset + 0] | source.Array[source.Offset + 1] << 8 | source.Array[source.Offset + 2] << 16);

                return len - 2 + HeadSize;

                // 客户端使用的是真实的包体长度，但是为什么传输时要 +2 ???
                // -2 是由于去除长度中带的消息Id长度，否则就和HeadSize重复了，坑死宝宝了。
                //var header = new Header(source);

                //// 断言
                //if (header.BodyLength > MaxPackageSize
                //    || header.MsgId >= 65536)
                //{
                //    ex = new Exception("Communication protocol is destroyed");
                //}

                //// 纠正在传输中服务器要求包长+2的问题
                //header.BodyLength -= 2;
                //header.Fill(source);

                //return header.BodyLength + Header.HeadSize;
            }
            catch (Exception e)
            {
                ex = e;
                return 0;
            }
        }

        /// <summary>
        /// 在字节流向上层传递之前，对字节流做的修饰处理
        /// </summary>
        /// <param name="source">需要处理的字节流</param>
        /// <param name="ex">用户自定义异常，如果产生异常那么这个数据包会被抛弃</param>
        /// <returns>被处理后的字节流</returns>
        public ArraySegment<byte> Receive(ArraySegment<byte> source, out Exception ex)
        {
            ex = null;

            //var len = (source.Array[source.Offset + 0] | source.Array[source.Offset + 1] << 8 | source.Array[source.Offset + 2] << 16);
            //var bcomp = (source.Array[source.Offset + 3] & 0x1) > 0;
            //var bencrypt = (source.Array[source.Offset + 3] & 0x2) > 0;

            //if (bencrypt)
            //{
            //    XorEncrypt.Decrypt(source.Array, source.Offset + HeadSize, len);
            //}

            //if (bcomp)
            //{
            //    var uncompdata = Snappy.Uncompress(source.Array, source.Offset + HeadSize, len);
            //}

            //try
            //{
            //    var header = new Header(source);

            //    // 进行解密
            //    if (header.Encrypt)
            //    {
            //        XorEncrypt.Decrypt(header.Body.Array, header.Body.Count, header.Body.Offset);
            //    }

            //    // 进行解压缩
            //    if (header.Compress)
            //    {
            //        // 不允许使用相同的数组进行压缩赋值，否则会导致bug
            //        // 这里使用了receiveCompressBuffer作为中转
            //        int unCompressCount;
            //        if (!Snappy.Uncompress(header.Body.Array, header.Body.Offset, header.Body.Count, receiveCompressBuffer,
            //            out unCompressCount))
            //        {
            //            ex = new Exception("message uncompress error.");
            //            return source;
            //        }

            //        header.BodyLength = unCompressCount;
            //        Buffer.BlockCopy(receiveCompressBuffer, 0, header.Body.Array, header.Body.Offset, unCompressCount);

            //        // 解压缩导致切片范围发生改变，所以我们重新定义切片范围 （包头 + 包体）
            //        source = new ArraySegment<byte>(source.Array, source.Offset, source.Offset + header.Body.Offset + header.BodyLength);
            //    }
            //}
            //catch (Exception exception)
            //{
            //    ex = exception;
            //}

            return source;
        }

        /// <summary>
        /// 在字节流向Socket传递之前对字节流做的修饰处理
        /// </summary>
        /// <param name="source">需要处理的字节流（包头 + 包体）</param>
        /// <param name="ex">用户自定义异常，如果产生异常那么这个数据包会被抛弃</param>
        /// <returns>被处理后的字节流</returns>
        public ArraySegment<byte> Send(ArraySegment<byte> source, out Exception ex)
        {
            ex = null;

            //try
            //{
            //    var header = new Header(source)
            //    {
            //        Compress = source.Count > MinCompressSize
            //    };

            //    // 压缩
            //    if (header.Compress)
            //    {
            //        var packerSize = Snappy.MaxCompressedLength(header.Body.Count) + Header.HeadSize;
            //        if (packerSize > header.Body.Array.Length)
            //        {
            //            ex = new Exception("Can not write in buffer");
            //            return source;
            //        }

            //        // 不允许使用相同的数组进行压缩赋值，否则会导致bug
            //        // 这里使用了sendCompressBuffer作为中转
            //        header.BodyLength = Snappy.Compress(header.Body.Array, header.Body.Offset, header.Body.Count, sendCompressBuffer, 0);
            //        Buffer.BlockCopy(sendCompressBuffer, 0, header.Body.Array, header.Body.Offset, header.BodyLength);

            //        // 由于压缩了导致切片发生变动所以我们重新定义切片范围
            //        header.Body = new ArraySegment<byte>(source.Array, header.Body.Offset, header.BodyLength);
            //    }

            //    // 加密
            //    if (header.Encrypt = UseEncrypt)
            //    {
            //        XorEncrypt.Encrypt(header.Body.Array, header.Body.Count, header.Body.Offset);
            //    }

            //    // 定义 包头 + 包体 的切片
            //    source = new ArraySegment<byte>(source.Array, source.Offset, source.Offset + header.Body.Offset + header.BodyLength);

            //    // 坑爹的 +2 不解释
            //    header.BodyLength += 2;

            //    // 将头信息填充到字节流中
            //    header.Fill(source);

            //    return source;
            //}
            //catch (Exception exception)
            //{
            //    ex = exception;
            //}

            return source;
        }
    }
}
