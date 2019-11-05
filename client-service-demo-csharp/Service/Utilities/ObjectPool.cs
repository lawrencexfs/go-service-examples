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
using System.Collections.Generic;
using GameBox;

namespace GameBox.Service
{
    internal class BaseObjectPool
    {
        public string typename = "";
        public List<object> objectList = new List<object>();
        public Func<object, object> createact;
        public BaseObjectPool()
        {
        }

        public object NewObj(Func<object, object> increateact, object args)
        {
            if (objectList.Count == 0)
            {
                createact = increateact;
                object objins = createact.Invoke(args);
                return objins;
            }
            else
            {
                int last = objectList.Count - 1;
                object data = objectList[last];
                objectList.RemoveAt(last);
                return data;
            }
        }

        public void Recycle(object data)
        {
            if (objectList.Contains(data))
                return;

            objectList.Add(data);
        }

        public void Clear()
        {
            objectList.Clear();
        }
    };

    public static class GlobalObjectPool<T> where T : new()
    {
        private static readonly Stack<T> pool = new Stack<T>();

        public static T Pop()
        {
            return pool.Count == 0 ? new T() : pool.Pop();
        }

        public static void Push(T obj)
        {
            Guard.NotNull(obj, "obj");
            pool.Push(obj);
        }
    }


    internal class GObjectPool
    {
        //全局唯一对象池
        protected static Dictionary<Type, BaseObjectPool> theObjectPool = new Dictionary<Type, BaseObjectPool>();

        public static T Create<T>() where T : new()
        {
            if (theObjectPool.ContainsKey(typeof(T)) == false)
            {
                theObjectPool.Add(typeof(T), new BaseObjectPool());
            }

            return (T)theObjectPool[typeof(T)].NewObj(x => new T(), null);
        }

        public static T Create<T>(Func<object, object> createfunc, object args)
        {
            if (theObjectPool.ContainsKey(typeof(T)) == false)
            {
                theObjectPool.Add(typeof(T), new BaseObjectPool());
            }

            return (T)theObjectPool[typeof(T)].NewObj(createfunc, args);
        }

        public static void Recycle(object ap)
        {
            if (theObjectPool.ContainsKey(ap.GetType()) == false)
            {
                theObjectPool.Add(ap.GetType(), new BaseObjectPool());
            }

            theObjectPool[ap.GetType()].Recycle(ap);
        }

        public static void Clear() //清理
        {
            foreach (var baseObjectPool in theObjectPool)
            {
                baseObjectPool.Value.Clear();
            }

        }

        public static int ReserveCount<T>()
        {
            if (theObjectPool.ContainsKey(typeof(T)) == false)
            {
                return 0;
            }

            return theObjectPool[typeof(T)].objectList.Count;
        }

    }
}