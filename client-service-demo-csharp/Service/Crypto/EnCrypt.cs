using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace GameBox.Service.Crypto
{
    internal class EnCrypt
    {
        private static readonly byte[] encryptKey = { 41, 247, 6, 255, 138, 78, 197, 129 };

        public static void Encrypt(byte[] msgBuf)
        {
            var k = 0;
            for (var i = 0; i < msgBuf.Length; i++)
            {
                var n = (byte)(i % 7 + 1); //移位长度
                var b = (byte)((byte)(msgBuf[i] << n) | (byte)((msgBuf[i] >> (8 - n)))); // 向左循环移位
                msgBuf[i] = (byte)(b ^ encryptKey[k]);
                k++;
                k = k % encryptKey.Length;
            }
        }

        public static void Encrypt(byte[] msgBuf, int offset, int count)
        {
            var k = 0;
            for (var i = offset; i < count + offset; i++)
            {
                var n = (byte)((i - offset) % 7 + 1); //移位长度
                var b = (byte)((byte)(msgBuf[i] << n) | (byte)((msgBuf[i] >> (8 - n)))); // 向左循环移位
                msgBuf[i] = (byte)(b ^ encryptKey[k]);
                k++;
                k = k % encryptKey.Length;
            }
        }
    }
}
