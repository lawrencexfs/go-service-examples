using GameBox.Pioneer;

namespace Kongfu
{
    public sealed class Context : Trait
    {
        public IWorld World { private get; set; }

        public IEntity CreateEntity()
        {
            return this.World.CreateEntity();
        }

        public IRoomManager Room { get; set; }

        public string IP { get; set; }

        public int Port { get; set; }
    }
}
