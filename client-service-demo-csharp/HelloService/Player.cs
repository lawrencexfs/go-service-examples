using GameBox.Service;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace HelloService
{
    public class Player : Entity
    {
        [GBoxRPC]
        private void Hello(string name, uint id)
        {
            Console.WriteLine("recv Hello msg from server, name:" + name + " id:" + id);
        }
    }
}
