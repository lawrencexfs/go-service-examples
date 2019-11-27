using GameBox.Pioneer;

namespace Kongfu
{
    public sealed class Message : Trait
    {
        public MessageDef Type { get; set; }

        public object[] Payload { get; set; }
    }
}
