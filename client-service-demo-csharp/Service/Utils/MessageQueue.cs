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

using System.Collections.Generic;

namespace GameBox.Service
{
#if MUTIL_THREAD_MESSAGEQUEUE
    public class MessageQueue<T>
    {
        List<T> empty = new List<T>();
        Queue<T> datas ;

        object lockobj = new object();

        public MessageQueue()
        {
            datas = new Queue<T>();
        }

        public void Enqueue(T data)
        {
            lock(lockobj)
            {
                datas.Enqueue(data);
            }
        }

        public List<T> GetQueue()
        {
            if (datas.Count == 0) return empty;

            List<T> l = null;

            lock(lockobj)
            {
                l = new List<T>( datas.Count );
                l.AddRange(datas);
                datas.Clear();
            }

            return l;
        }

        public void Clear()
        {
            datas.Clear();
        }

        public T Dequeue()
        {
            T r;
            lock (lockobj)
            {
                if (datas.Count == 0)
                   return default(T);

                r = datas.Dequeue();
            }

             return r;
         }
    }
#else
    internal class MessageQueue<T>
    {
        private bool _lock = false;
        private List<T> _tmp = new List<T>();
        private Queue<T> datas = new Queue<T>();

        public int Count
        {
            get { return this.datas.Count; }
        }

        public void Enqueue(T data)
        {
            if (this._lock)
                this._tmp.Add(data);
            else
                this.datas.Enqueue(data);
        }

        public Queue<T> GetQueue()
        {
            this._lock = true;
            return datas;
        }

        public T Dequeue()
        {
            if (this.datas.Count > 0)
                return this.datas.Dequeue();
            return default(T);
        }

        public void Clear()
        {
            this.datas.Clear();
            this._lock = false;
            foreach (var obj in this._tmp)
                this.Enqueue(obj);
            this._tmp.Clear();
        }
    }
#endif
}