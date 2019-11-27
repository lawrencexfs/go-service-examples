using ProtoBuf;
using System.IO;

namespace Kongfu
{
    public sealed class Proto : IProto
    {
        public Proto(object obj)
        {
            this.obj = obj;
        }

        public Proto(byte[] bytes)
        {
            this.bytes = bytes;
        }

        public T ToObject<T>()
        {
            return Deserialize<T>(bytes);
        }

        public byte[] ToBytes()
        {
            return Serialize(obj);
        }

        private byte[] Serialize<T>(T obj)
        {
            using (var steam = new MemoryStream())
            {
                Serializer.Serialize(steam, obj);
                return steam.ToArray();
            }
        }

        private T Deserialize<T>(byte[] bytes)
        {
            using (var stream = new MemoryStream(bytes))
            {
                var obj = default(T);
                obj = Serializer.Deserialize<T>(stream);
                return obj;
            }
        }

        private readonly byte[] bytes;
        private readonly object obj;
    }
}