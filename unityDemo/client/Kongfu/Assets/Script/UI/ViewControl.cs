using GameBox.Pioneer;
using UnityEngine;

namespace Kongfu
{
    public abstract class ViewControl : Control
    {
        public override void Dispose()
        {
            OnUnload();
        }

        public override void OnInit(ITraitContainer container)
        {
            this.container = container;

            var m = container.NewMatcher();
            m.HasTrait<Window>().HasTag(Window.Inited);
            this.filter = container.GetFilter(this, TupleType.Job, m);
        }

        public override void OnUpdate(ITraitContainer container, float deltaTime)
        {
            if (null != this.filter.Target)
            {
                var e = this.filter.Target;
                var w = e.GetTrait<Window>();

                if (!this.inited)
                {
                    this.inited = true;
                    OnLoad();
                }

                OnUpdate(deltaTime);

                if (-1 != w.FadeOutTimer && w.FadeOutTimer > 0)
                {
                    w.FadeOutTimer -= Time.deltaTime;
                    if (w.FadeOutTimer <= 0)
                    {
                        w.FadeOutTimer = -1;
                        Game.CloseWindow(w.Path);
                    }
                }
            }
        }

        protected virtual void OnLoad()
        {
        }

        protected virtual void OnUnload()
        {
        }

        protected virtual void OnUpdate(float deltaTime)
        {
        }

        protected void Close()
        {
            var e = this.filter.Target;
            var w = e.GetTrait<Window>();
            Game.CloseWindow(w.Path);
        }

        protected Transform Find(string name)
        {
            var window = this.container.GetTrait<Window>();
            return window.View.transform.Find(name);
        }

        private ITraitContainer container = null;
        private IEntityFilter filter = null;
        private bool inited = false;
    }
}
