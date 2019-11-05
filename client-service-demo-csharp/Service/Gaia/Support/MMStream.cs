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
using System.IO;

namespace GameBox.Service
{
    internal class MMStream
    {
        public byte[] Buf;
        public int WPos;
        public int RPos;
        public int Capicity;

        public MMStream()
        {
            this.Reset(null);
        }
        public MMStream(byte[] bytes)
        {
            this.Reset(bytes);
        }

        public void Reset()
        {
            RPos = 0;
            WPos = 0;
        }

        public void Reset(byte[] bytes,int inLen=-1)
        {
            Buf = bytes;
            RPos = 0;
            WPos = 0;
            if (bytes != null)
                Capicity = (inLen==-1)? bytes.Length: inLen;
            else
                Capicity = 0;
        }
		public void Reset(byte[] bytes, int offset, int count)
        {
            Buf = bytes;
            RPos = offset;
            WPos = offset;
            Capicity = count + offset;
        }
        public void ResetByFixByteArray(FixedLengthArray<byte> fixbytes)
        {
            Reset(fixbytes.getArray(), fixbytes.usedLength);
        }

        public void WriteBool(bool b)
        {
            if(!IsWEnough(1))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buf[WPos] = (byte)(b ? 1 : 0);
            WPos += 1;
        }
        public void WriteByte(byte b1)
        {
            if (!IsWEnough(1))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buf[WPos] = b1;
            WPos += 1;
        }

        public void WriteBytes(byte[] p)
        {
            if (p == null)
            {
                WriteUInt16(0);
                return;
            }

            int len = p.Length;
            WriteUInt16((ushort)len);
            WriteBytes(p, len);
        }

        public void WriteBuf(byte[] p , int len)
        {
            if (p == null)
            {
                WriteUInt16(0);
                return;
            }

            WriteUInt16((ushort)len);
            WriteBytes(p, len);
        }

        public void WriteBytes(byte[] p, int len)
        {
            if (p == null)
            {
                WriteUInt16(0);
                return;
            }

            if (len == 0)
            {
                return;
            }

            if (!IsWEnough(len))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buffer.BlockCopy(p, 0, Buf, WPos, len);
            WPos += len;
        }

        public void WriteInt16(short n)
        {
            if (!IsWEnough(2))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buf[WPos + 0] = (byte)n;
            Buf[WPos + 1] = (byte)(n >> 8);
            WPos += 2;
        }

        public void WriteUInt16(ushort n)
        {
            if (!IsWEnough(2))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buf[WPos + 0] = (byte)n;
            Buf[WPos + 1] = (byte)(n >> 8);
            WPos += 2;
        }

        public void WriteInt32(int n)
        {
            if (!IsWEnough(4))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buf[WPos + 0] = (byte)(n & 0xff);
            Buf[WPos + 1] = (byte)((n >> 8) & 0xff);
            Buf[WPos + 2] = (byte)((n >> 16) & 0xff);
            Buf[WPos + 3] = (byte)((n >> 24) & 0xff);

            WPos += 4;
        }

        public void WriteUInt32(uint n)
        {
            if (!IsWEnough(4))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buf[WPos + 0] = (byte)(n & 0xff);
            Buf[WPos + 1] = (byte)((n >> 8) & 0xff);
            Buf[WPos + 2] = (byte)((n >> 16) & 0xff);
            Buf[WPos + 3] = (byte)((n >> 24) & 0xff);

            WPos += 4;
        }

        public void WriteInt64(long n)
        {
            WriteBytes(BitConverter.GetBytes(n), 8);
        }

        public void WriteUInt64(UInt64 n)
        {
            WriteBytes(BitConverter.GetBytes(n) , 8);
        }

        public void WriteFloat(float n)
        {
            WriteBytes(BitConverter.GetBytes(n) , 4);
        }

        public void WriteDouble(double n)
        {
            WriteBytes(BitConverter.GetBytes(n) , 8);
        }

        public void WriteString(string str)
        {
            byte[] bytes = System.Text.Encoding.UTF8.GetBytes(str);
            WriteBytes(bytes);
        }

        public byte ReadByte()
        {
            if (!IsREnough(1))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            RPos++;
            return Buf[RPos - 1];
        }


        public bool ReadBoolean()
        {
            if (!IsREnough(1))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            RPos++;
            return Convert.ToBoolean(Buf[RPos - 1]);
        }

        public byte[] ReadBytes(int count)
        {
            if(!IsREnough(count))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            byte[] bytes = new byte[count];
            Buffer.BlockCopy(Buf, RPos, bytes, 0, count);
            RPos += count;
            return bytes;
        }
        

        public int ReadFixByteArray(ref FixedLengthArray<byte> data)
        {
            int dataActLen = ReadUInt16();

            if (data.totalLength < dataActLen)
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            data.Clear();
            data.SetUsedLength(dataActLen);
            var arr = data.getArray();
            ReadBytes(ref arr, dataActLen);
            return dataActLen + 2;
        }

        public int ReadBytes(ref byte[] outputStream, int count)
        {
            if (!IsREnough(count))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            if (outputStream.Length < count)
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            Buffer.BlockCopy(Buf, RPos, outputStream, 0, count);
            RPos += count;
            return count;
        }

        public byte[] ReadBytes()
        {
            var len = ReadUInt16();
            return ReadBytes(len);
        }

        public char ReadChar()
        {
            return (char)ReadByte();
        }

        public float ReadFloat()
        {
            //byte[] bytes = ReadBytes(4);
            if (!IsREnough(4))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            float ret =  BitConverter.ToSingle(Buf, RPos);
            RPos += 4;
            return ret;
        }

        public float Read16BitFloat()
        {
            float bytes = (float)ReadInt16();

            bytes = bytes / 100.0f;
            return bytes;
        }

        public double ReadDouble()
        {
            //byte[] bytes = ReadBytes(8);
            if (!IsREnough(8))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }
            double ret =BitConverter.ToDouble(Buf, RPos );
            RPos += 8;
            return ret;
        }

        public short ReadInt16()
        {
            //int a = (int)ReadByte();
            //int b = (int)ReadByte();
            //return (short)(a | (b << 8));

            if (!IsREnough(2))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }
            short ret = BitConverter.ToInt16(Buf, RPos );
            RPos += 2;
            return ret;
        }
        public ushort ReadUInt16()
        {
            //int a = (int)ReadByte();
            //int b = (int)ReadByte();
            //return (ushort)((b << 8) | a);
            if (!IsREnough(2))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }
            ushort ret = BitConverter.ToUInt16(Buf, RPos );
            RPos += 2;
            return ret;
        }

        public int ReadInt32()
        {
            if (!IsREnough(4))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }
            int ret = BitConverter.ToInt32(Buf, RPos );
            RPos += 4;
            return ret;

            /*
            int c1 = (int)ReadByte();
            int c2 = (int)ReadByte();
            int c3 = (int)ReadByte();
            int c4 = (int)ReadByte();
            return (c4 << 24) | (c3 << 16) | (c2 << 8) | c1;
            */
        }
        /*
        public int ReadInt24()
        {
            int c1 = (int)ReadByte();
            int c2 = (int)ReadByte();
            int c3 = (int)ReadByte();
            return c1 | (c2 << 8) | (c3 << 16);

        }
        */
        public uint ReadUInt32()
        {
            if (!IsREnough(4))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }
            uint ret = BitConverter.ToUInt32(Buf, RPos );
            RPos += 4;
            return ret;

            /*
            uint c1 = (uint)ReadByte();
            uint c2 = (uint)ReadByte();
            uint c3 = (uint)ReadByte();
            uint c4 = (uint)ReadByte();
            return (c4 << 24) | (c3 << 16) | (c2 << 8) | c1;
            */
        }
        public long ReadInt64()
        {
            if (!IsREnough(8))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }
            long ret = BitConverter.ToInt64(Buf, RPos );
            RPos += 8;
            return ret;

            /*         long c1 = (long)ReadByte();
                       long c2 = (long)ReadByte();
                       long c3 = (long)ReadByte();
                       long c4 = (long)ReadByte();
                       long c5 = (long)ReadByte();
                       long c6 = (long)ReadByte();
                       long c7 = (long)ReadByte();
                       long c8 = (long)ReadByte();

                       return (c8 << 56) | (c7 << 48) | (c6 << 40) | (c5 << 32) | (c4 << 24) | (c3 << 16) | (c2 << 8) | c1;
            */
        }
        public ulong ReadUInt64()
        {
            if (!IsREnough(8))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            ulong ret = BitConverter.ToUInt64(Buf, RPos );
            RPos += 8;
            return ret;
            
            /*
            //byte[] bytes = ReadBytes(8);
            ulong c1 = ReadByte();      //(ulong)bytes[0];
            ulong c2 = ReadByte();      //(ulong)bytes[1];
            ulong c3 = ReadByte();      //(ulong)bytes[2];
            ulong c4 = ReadByte();      //(ulong)bytes[3];
            ulong c5 = ReadByte();      //(ulong)bytes[4];
            ulong c6 = ReadByte();      //(ulong)bytes[5];
            ulong c7 = ReadByte();      //(ulong)bytes[6];
            ulong c8 = ReadByte();      //(ulong)bytes[7];
            return (c8 << 56) | (c7 << 48) | (c6 << 40) | (c5 << 32) | (c4 << 24) | (c3 << 16) | (c2 << 8) | c1;
            */
        }

        public string ReadString()
        {
            int len = (int)ReadUInt16();

            if (!IsREnough(len))
            {
                throw new EndOfStreamException(string.Format("Capicity={0}, WPos={1}, RPos={2}", Capicity, WPos, RPos));
            }

            //byte[] bytes = ReadBytes(len);
            //return System.Text.Encoding.UTF8.GetString(bytes);
            string ret = System.Text.Encoding.UTF8.GetString(Buf,RPos ,len);
            RPos += len;

            return ret;
        }
        
        public bool IsWEnough(int size)
        {
            return (Capicity - WPos) >= size;
        }

        public bool IsREnough(int size)
        {
            return (Capicity - RPos) >= size;
        }

        public bool IsEOF()
        {
            return RPos == Capicity;
        }
    }
}

