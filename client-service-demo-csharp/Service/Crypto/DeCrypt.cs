using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace GameBox.Service.Crypto
{
    internal class DeCrypt
    {
        private static readonly byte[] decryptkey = { 253, 1, 56, 52, 62, 176, 42, 138 };

        public static void Decrypt(byte[] msgBuf)
        {
            var k = 0;
            for (var i = 0; i < msgBuf.Length; i++)
            {
                var b = (byte)(msgBuf[i] ^ decryptkey[k]);
                var n = (byte)(i % 7 + 1); //移位长度
                msgBuf[i] = (byte)((byte)(b >> n) | (byte)(b << (8 - n))); // 向右循环移位

                k++;
                k = k % decryptkey.Length;
            }
        }

        public static void Decrypt(byte[] msgBuf, int offset, int count)
        {
            var k = 0;
            for (var i = offset; i < count + offset; i++)
            {
                var b = (byte)(msgBuf[i] ^ decryptkey[k]);
                var n = (byte)((i - offset) % 7 + 1); //移位长度
                msgBuf[i] = (byte)((byte)(b >> n) | (byte)(b << (8 - n))); // 向右循环移位

                k++;
                k = k % decryptkey.Length;
            }
        }
    }
}
