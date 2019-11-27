using GameBox.Pioneer;
using usercmd;

namespace Kongfu
{
    public sealed class Item : Trait
    {
        public const string Inited = "item.inited";

        public MsgBall Data { get; set; }
    }
}