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

namespace GameBox.Service
{
    internal class XorEncrypt
    {
        private static readonly byte[] encryptKey = { 41, 247, 6, 255, 138, 78, 197, 129 };
        private static readonly byte[] decryptkey = { 253, 1, 56, 52, 62, 176, 42, 138 };

        //public static void Encrypt(byte[] msgBuf, int bufLen, int offset = 0)
        //{
        //    var k = 0;
        //    for (var i = offset; i < bufLen + offset; i++)
        //    {
        //        var n = (byte)((i - offset) % 7 + 1); //移位长度
        //        var b = (byte)((byte)(msgBuf[i] << n) | (byte)((msgBuf[i] >> (8 - n)))); // 向左循环移位
        //        msgBuf[i] = (byte)(b ^ encryptKey[k]);
        //        k++;
        //        k = k % encryptKey.Length;
        //    }
        //}

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

        //public static void Decrypt(byte[] msgBuf, int bufLen, int offset = 0)
        //{
        //    var k = 0;
        //    for (var i = offset; i < bufLen + offset; i++)
        //    {
        //        var b = (byte)(msgBuf[i] ^ decryptkey[k]);
        //        var n = (byte)((i - offset) % 7 + 1); //移位长度
        //        msgBuf[i] = (byte)((byte)(b >> n) | (byte)(b << (8 - n))); // 向右循环移位

        //        k++;
        //        k = k % decryptkey.Length;
        //    }
        //}

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