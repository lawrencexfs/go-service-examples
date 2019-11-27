using System;
using System.Collections;
using System.Collections.Generic;

namespace Kongfu
{
    public sealed class LazySet<T> : IEnumerable<T>
    {
        public bool Add(T item)
        {
            if (this.locked)
            {
                this.actions.Enqueue(new KeyValuePair<T, bool>(item, true));
                return true;
            }
            else
            {
                return this.set.Add(item);
            }
        }

        public bool Remove(T item)
        {
            if (this.locked)
            {
                this.actions.Enqueue(new KeyValuePair<T, bool>(item, false));
                return true;
            }
            else
            {
                return this.set.Remove(item);
            }
        }

        public bool Contains(T item)
        {
            return this.set.Contains(item);
        }

        public int Count
        {
            get
            {
                return this.set.Count;
            }
        }

        public IEnumerator<T> GetEnumerator()
        {
            return new Enumerator(this);
        }

        IEnumerator IEnumerable.GetEnumerator()
        {
            return new Enumerator(this);
        }

        private void ExecuteActions()
        {
            if (this.locked)
            {
                return;
            }

            while (this.actions.Count > 0)
            {
                var action = this.actions.Dequeue();
                if (action.Value)
                {
                    this.set.Add(action.Key);
                }
                else
                {
                    this.set.Remove(action.Key);
                }
            }
        }

        private struct Enumerator : IEnumerator<T>
        {
            public Enumerator(LazySet<T> lazySet)
            {
                this.lazySet = lazySet;
                this.it = lazySet.set.GetEnumerator();
                LockSet();
            }

            public T Current
            {
                get
                {
                    return this.it.Current;
                }
            }

            object IEnumerator.Current
            {
                get
                {
                    return this.it.Current;
                }
            }

            public void Dispose()
            {
                UnlockSet();
            }

            public bool MoveNext()
            {
                var next = this.it.MoveNext();
                if (!next)
                {
                    UnlockSet();
                }

                return next;
            }

            public void Reset()
            {
                throw new NotSupportedException();
            }

            private void LockSet()
            {
                this.lazySet.locked = true;
            }

            private void UnlockSet()
            {
                this.lazySet.locked = false;
                this.lazySet.ExecuteActions();
            }

            private LazySet<T> lazySet;
            private IEnumerator<T> it;
        }

        private HashSet<T> set = new HashSet<T>();
        private bool locked = false;
        private Queue<KeyValuePair<T, bool>> actions = new Queue<KeyValuePair<T, bool>>();
    }
}
