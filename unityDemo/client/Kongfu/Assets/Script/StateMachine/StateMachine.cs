using GameBox.Pioneer;

namespace Kongfu
{
    public sealed class StateMachine : Control
    {
        public void ChangeState(State state)
        {
            if (null != this.state)
            {
                this.state.OnExit(this);
            }

            this.state = state;

            if (null != this.state)
            {
                this.state.OnEnter(this);
            }
        }

        public override void OnInit(ITraitContainer container)
        {
            this.owner = container;
        }

        public override void OnUpdate(ITraitContainer container, float deltaTime)
        {
            if (null != this.state)
            {
                this.state.OnExecute(this);
            }
        }

        public override void Dispose()
        {
            if (null != this.state)
            {
                this.state.OnExit(this);
                this.state = null;
            }
        }

        public T GetTrait<T>() where T : Trait
        {
            return this.owner.GetTrait<T>();
        }

        private ITraitContainer owner = null;
        private State state = null;
    }
}

