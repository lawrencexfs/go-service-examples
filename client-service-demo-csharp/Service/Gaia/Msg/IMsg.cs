/*
* This file is part of the GameBox package.
*
* (c) Giant - MouGuangYi<mouguangyi@ztgame.com> , tanxiaoliang<tanxiaoliang@ztgame.com>
*
* For the full copyright and license information, please view the LICENSE
* file that was distributed with this source code.
*
* Document: http://192.168.150.238/GameBox/help/ 
*/

namespace GameBox.Service
{
    internal interface IMsg
    {
        void Marshal(MMStream stream);
        void UnMarshal(MMStream stream);

        void Recycle();     //回收接口
    }
}
