namespace Kongfu
{
    public interface IProto
    {
        T ToObject<T>();
        byte[] ToBytes();
    }
}
