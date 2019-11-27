using GameBox.Pioneer;
using System;
using System.Collections.Generic;

namespace Kongfu
{
    public sealed class MessageSystem : GameBox.Pioneer.System, IMessageObserver
    {
        public override void OnInit(IEntityContainer container)
        {
            var m = container.NewMatcher();
            m.HasTrait<Message>();
            this.filter = container.GetFilter(this, TupleType.Job, m);
        }

        public override void OnUpdate(IEntityContainer container, float deltaTime)
        {
            foreach (var e in this.filter.Target)
            {
                var msg = e.GetTrait<Message>();
                LazySet<Action<MessageDef, object[]>> set = null;
                if (this.handlers.TryGetValue(msg.Type, out set) && set.Count > 0)
                {
                    foreach (var handler in set)
                    {
                        handler(msg.Type, msg.Payload);
                    }
                }

                e.Dispose();
            }
        }

        public bool On(MessageDef type, Action<MessageDef, object[]> handler)
        {
            if (null == handler)
            {
                return false;
            }

            LazySet<Action<MessageDef, object[]>> set = null;
            if (!this.handlers.TryGetValue(type, out set))
            {
                this.handlers.Add(type, set = new LazySet<Action<MessageDef, object[]>>());
            }

            return set.Add(handler);
        }

        public bool Off(MessageDef type, Action<MessageDef, object[]> handler)
        {
            if (null == handler)
            {
                return false;
            }

            LazySet<Action<MessageDef, object[]>> set = null;
            if (!this.handlers.TryGetValue(type, out set))
            {
                return false;
            }

            return set.Remove(handler);
        }

        private IEntitiesFilter filter = null;
        private Dictionary<MessageDef, LazySet<Action<MessageDef, object[]>>> handlers = new Dictionary<MessageDef, LazySet<Action<MessageDef, object[]>>>();
    }
}
