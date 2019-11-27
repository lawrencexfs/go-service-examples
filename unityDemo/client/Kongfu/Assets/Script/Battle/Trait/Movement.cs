using GameBox.Pioneer;
using usercmd;

namespace Kongfu
{
    public sealed class Movement : Trait
    {
        public const string Inited = "movement.inited";

        public BallMove Move
        {
            get
            {
                return this.move;
            }
            set
            {
                this.lastMove = this.move;
                this.move = value;
            }
        }

        public BallMove LastMove
        {
            get
            {
                return lastMove;
            }
        }

        private BallMove move = null;
        private BallMove lastMove = null;
    }
}