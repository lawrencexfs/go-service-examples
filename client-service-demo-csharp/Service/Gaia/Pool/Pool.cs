using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace GameBox.Service
{
    internal partial class Pool<T>
    {
        object lockObj = new object();
        Stack<T> objPool = new Stack<T>();

        internal T Get()
        {
            lock (lockObj)
            {
                if (objPool.Count > 0)
                {
                    return objPool.Pop();
                }

                T obj = default(T);
                objPool.Push(obj);
                return obj;
            }
        }

        internal void Put(T obj)
        {
            lock (lockObj)
            {
                objPool.Push(obj);
            }
        }
    }

    internal class MMStreamPool
    {
        static int[] sizes = new int[] {
            1 << 4,
            1 << 5,
            1 << 6,
            1 << 7,
            1 << 8,
            1 << 9,
            1 << 10,
            1 << 11,
            1 << 12,
            1 << 13,
            1 << 14,
            1 << 15,
            1 << 16,
        };

        const int MaxSize = 1 << 16;

        static Pool<MMStream>[] pools;
        static Pool<MMStream> pool;

        internal static void Init()
        {
            pools = new Pool<MMStream>[sizes.Length];

            for(int i = 0; i < sizes.Length; i++)
            {
                pools[i] = new Pool<MMStream>();
            }

            pool = new Pool<MMStream>();
        }

        internal static MMStream Get(int size = 65536)
        {
            if(size <= MaxSize)
            {
                for (int i = 0; i < sizes.Length; i++)
                {
                    if(size <= sizes[i])
                    {
                        var obj = pools[i].Get();
                        if(obj != null)
                        {
                            if (obj.Buf == null)
                            {
                                obj.Reset(new byte[sizes[i]]);
                            }
                        }
                        else
                        {
                            obj = new MMStream(new byte[sizes[i]]);
                        }
                        
                        return obj;
                    }
                }
            }

            return new MMStream(new byte[size]);
        }

        internal static void Put(MMStream obj)
        {
            var size = obj.Buf.Length;
            if (size < MaxSize)
            {
                for (int i = 0; i < sizes.Length; i++)
                {
                    if (size == sizes[i])
                    {
                        pools[i].Put(obj);
                    }
                }
            }
        }

        internal static MMStream GetTmp()
        {
            var obj = pool.Get();
            if (obj == null)
                obj = new MMStream();
            return obj;
        }

        internal static void PutTmp(MMStream obj)
        {
            obj.Reset(null);
            pool.Put(obj);
        }
    }
}
