using GameBox.Channel;
using System;

namespace Kongfu
{
    public sealed class Fragment : IFragment
    {
        public int MaxPackageSize
        {
            get
            {
                return 65536;
            }
        }

        public int Input(ArraySegment<byte> source, out Exception ex)
        {
            ex = null;
            try
            {
                if (source.Count < HeaderLength)
                {
                    return 0;
                }

                var len1 = source.Array[source.Offset + 0];
                var len2 = source.Array[source.Offset + 1];
                var len3 = source.Array[source.Offset + 2];
                var bodyLength = len1 | (len2 << 8) | (len3 << 16);
                var packageLength = HeaderLength + bodyLength;
                return source.Count < packageLength ? 0 : packageLength;
            }
            catch (Exception e)
            {
                ex = e;
                return 0;
            }
        }

        public ArraySegment<byte> Receive(ArraySegment<byte> source, out Exception ex)
        {
            ex = null;
            return source;
        }

        public ArraySegment<byte> Send(ArraySegment<byte> source, out Exception ex)
        {
            ex = null;
            return source;
        }

        public const int HeaderLength = 4;
        public const int CmdIdLength = 2;
    }
}
