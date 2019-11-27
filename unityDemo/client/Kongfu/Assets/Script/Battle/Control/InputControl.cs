using GameBox.Pioneer;
using UnityEngine;
using usercmd;

namespace Kongfu
{
    public sealed class InputControl : Control
    {
        public BattleManager Manager
        {
            private get; set;
        }

        public override void Dispose()
        {
            Game.Off(MessageDef.JoyStickDragging, OnJoyStickDragging);
            Game.Off(MessageDef.JoyStickStop, OnJoyStickStop);
            Game.Off(MessageDef.Attack, OnAttack);
        }

        public override void OnInit(ITraitContainer container)
        {
            Game.On(MessageDef.JoyStickDragging, OnJoyStickDragging);
            Game.On(MessageDef.JoyStickStop, OnJoyStickStop);
            Game.On(MessageDef.Attack, OnAttack);
        }

        private void OnJoyStickDragging(MessageDef type, object[] payload)
        {
            var x = (float)payload[0];
            var y = (float)payload[1];
            var angle = (int)Vector2.Angle(Vector2.right, new Vector2(x, y));
            SendJoyStick(y > 0 ? 360 - angle : angle);
        }

        private void OnJoyStickStop(MessageDef type, object[] payload)
        {
            SendJoyStick(0, 0);
        }

        private void OnAttack(MessageDef type, object[] payload)
        {
            var msg = new MsgCastSkill();
            msg.skillid = 100;
            this.Manager.SendPacket((int)MsgTypeCmd.CastSkill, msg);
        }

        private void SendJoyStick(int angle, int power = 100)
        {
            uint face = this.Manager.GetClosetPlayerBallId(3);
            if (lastAngle == angle && lastFace == face)
            {
                return;
            }

            lastAngle = angle;
            lastFace = face;

            moveMsg.angle = angle;
            moveMsg.power = power;
            moveMsg.face = face;
            this.Manager.SendPacket((int)MsgTypeCmd.Move, moveMsg);
        }

        private MsgMove moveMsg = new MsgMove();
        private int lastAngle;
        private uint lastFace;
    }
}