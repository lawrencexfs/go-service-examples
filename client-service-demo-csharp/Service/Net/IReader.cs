using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace GameBox.Service.Net
{
    public interface IReader
    {
        int ReadFull(byte[] buffer);
        int ReadFull(byte[] buffer, int offset, int count);

        byte ReadByte();
        sbyte ReaeSByte();
        short ReadShort();
        ushort ReadUShort();
        int ReadInt();
        uint ReadUInt();
        long ReadLong();
        ulong ReadULong();
    }
}
