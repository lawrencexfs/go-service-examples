using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Text;

namespace GameBox.Service.Net
{
    //small endian reader
    internal class TcpSocketReader : IReader
    {
        private System.Net.Sockets.Socket socket;
        private byte[] buffer = new byte[8];

        public TcpSocketReader(System.Net.Sockets.Socket socket)
        {
            this.socket = socket;
        }

        public int ReadFull(byte[] buff)
        {
            try
            {
                var rlen = this.socket.Receive(buff);
                while (rlen < buff.Length)
                {
                    rlen += this.socket.Receive(buff, rlen, buff.Length - rlen, SocketFlags.None);
                }
                return rlen;
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public int ReadFull(byte[] buff, int offset, int count)
        {
            try
            {
                var rlen = this.socket.Receive(buff, offset, count, SocketFlags.None);
                while (rlen < count)
                {
                    rlen += this.socket.Receive(buff, offset + rlen, count - rlen, SocketFlags.None);
                }
                return rlen;
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public byte ReadByte()
        {
            try
            {
                ReadFull(buffer, 0, 1);
                return buffer[0];
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public int ReadInt()
        {
            try
            {
                ReadFull(buffer, 0, 4);
                return buffer[0] | (buffer[1] << 8) | (buffer[2] << 16) | (buffer[3] << 24);
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public long ReadLong()
        {
            try
            {
                ReadFull(buffer, 0, 8);
                return (long)buffer[0] | ((long)buffer[1] << 8) | ((long)buffer[2] << 16) | ((long)buffer[3] << 24) | ((long)buffer[4] << 32) | ((long)buffer[5] << 40) | ((long)buffer[6] << 48) | ((long)buffer[7] << 56);
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public short ReadShort()
        {
            try
            {
                ReadFull(buffer, 0, 2);
                return (short)(buffer[0] | (buffer[1] << 8));
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public uint ReadUInt()
        {
            try
            {
                ReadFull(buffer, 0, 4);
                return buffer[0] | ((uint)buffer[1] << 8) | ((uint)buffer[2] << 16) | ((uint)buffer[3] << 24);
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public ulong ReadULong()
        {
            try
            {
                ReadFull(buffer, 0, 8);
                return (ulong)buffer[0] | ((ulong)buffer[1] << 8) | ((ulong)buffer[2] << 16) | ((ulong)buffer[3] << 24) | ((ulong)buffer[4] << 32) | ((ulong)buffer[5] << 40) | ((ulong)buffer[6] << 48) | ((ulong)buffer[7] << 56);
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public ushort ReadUShort()
        {
            try
            {
                ReadFull(buffer, 0, 2);
                return (ushort)(buffer[0] | (buffer[1] << 8));
            }
            catch (Exception e)
            {
                throw e;
            }
        }

        public sbyte ReaeSByte()
        {
            try
            {
                ReadFull(buffer, 0, 1);
                return (sbyte)buffer[0];
            }
            catch (Exception e)
            {
                throw e;
            }
        }
    }
}
