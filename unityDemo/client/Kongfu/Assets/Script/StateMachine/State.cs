namespace Kongfu
{
    public abstract class State
    {
        public virtual void OnEnter(StateMachine stateMachine)
        {
        }

        public virtual void OnExecute(StateMachine stateMachine)
        {
        }

        public virtual void OnExit(StateMachine stateMachine)
        {
        }
    }
}
