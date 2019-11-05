using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace GameBox.Service.Net
{
    public interface IParser
    {
        void StartParse(IReader reader);
    }
}
