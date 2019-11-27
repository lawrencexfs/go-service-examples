using GameBox.Pioneer;
using usercmd;

namespace Kongfu
{
    public sealed class Player : Trait
    {
        public const string Inited = "player.inited";
        public const string MainPlayer = "player.main";

        public MsgPlayer Data { get; set; }

        public bool IsDead { get; set; }
    }
}
