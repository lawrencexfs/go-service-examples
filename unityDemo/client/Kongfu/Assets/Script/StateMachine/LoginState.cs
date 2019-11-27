using System;
using UnityEngine;

namespace Kongfu
{
    public class LoginState : State
    {
        public override void OnEnter(StateMachine stateMachine)
        {
            Debug.Log("Enter Login State.");

            this.action = (message, payload) =>
            {
                switch (message)
                {
                case MessageDef.Login:
                    {
                        var ip = (string)payload[0];
                        var port = (int)payload[1];
                        var ctx = stateMachine.GetTrait<Context>();
                        ctx.IP = ip;
                        ctx.Port = port;

                        Game.OpenWindow("UI/Loading");
                        stateMachine.ChangeState(new BattleState());
                    }
                    break;
                }
            };

            Game.On(MessageDef.Login, this.action);
            Game.OpenWindow<LoginControl>("UI/Login");
        }

        public override void OnExit(StateMachine stateMachine)
        {
            Game.CloseWindow("UI/Login");
            Game.Off(MessageDef.Login, this.action);
            this.action = null;
        }

        private Action<MessageDef, object[]> action = null;
    }
}
