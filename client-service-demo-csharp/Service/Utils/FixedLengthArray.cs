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

namespace GameBox.Service
{
    /// <summary>
    /// 定长数组
    /// 为了消息数据的优化
    /// </summary>
    internal class FixedLengthArray<T>
    {
        public int totalLength { get { return array.Length; } }
        public int usedLength { get; private set; }
        public T this[int index]
        {
            get
            {
                if (index >= usedLength)
                {
                    return default(T);
                }
                return array[index];
            }
        }
        readonly static T defaultT = default(T);
        T[] array;
        bool enableGetArray;

        public FixedLengthArray(int initLength, bool enableGetArray = true)
        {
            array = new T[initLength];
            usedLength = 0;
            this.enableGetArray = enableGetArray;
        }
        //Resize 可以调整长度，但是将清除内容
        public void Resize(int initLength)
        {
            array = new T[initLength];
            usedLength = 0;
        }

        public T[] getArray()
        {
            if (!enableGetArray)
            {
                throw new MethodAccessException("没有权限获取array");
            }
            return array;
        }

        public int Add(T t)
        {
            if (usedLength >= totalLength)
            {
                // throw new IndexOutOfRangeException("usedLength >= totalLength.  usedLength:" + usedLength + "    totalLength:" + totalLength);
                return -1;
            }
            array[usedLength++] = t;
            return usedLength;
        }

        /// <summary>
        /// </summary>
        /// <param name="t"></param>
        /// <param name="length">-1表示全部</param>
        /// <returns></returns>
        public int Set(T[] t, int length = -1)
        {
            if (t == null || t.Length == 0 || length == 0)
            {
                Clear();
                return 0;
            }

            var setLength = Math.Min(length == -1 ? t.Length : length, totalLength);
            for (int i = 0; i < setLength; i++)
            {
                array[i] = t[i];
            }
            if (setLength < usedLength)
            {
                for (int i = setLength; i < usedLength; i++)
                {
                    array[i] = defaultT;
                }
            }
            usedLength = setLength;
            return usedLength;
        }

        public void Clear()
        {
            for (int i = 0; i < usedLength; i++)
            {
                array[i] = defaultT;
            }
            usedLength = 0;
        }

        public void SetUsedLength(int length)
        {
            if (length < 0)
            {
                usedLength = 0;
            }
            else if (length > totalLength)
            {
                usedLength = totalLength;
            }
            else
            {
                usedLength = length;
            }
        }

        public void ForEach(Action<T> action)
        {
            if (action == null)
            {
                return;
            }
            for (int i = 0; i < usedLength; i++)
            {
                action(array[i]);
            }
        }

        public bool Contains(T t)
        {
            for (int i = 0; i < usedLength; i++)
            {
                if (array[i].Equals(t)) return true;
            }
            return false;
        }
    }
}