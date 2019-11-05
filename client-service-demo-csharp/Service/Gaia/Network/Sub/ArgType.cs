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
    /// <summary>
    /// 参数类型
    /// </summary>
    internal enum ArgType
    {
        typeUint8 = 1,
        typeUint16 = 2,
        typeUint32 = 3,
        typeUint64 = 4,
        typeInt8 = 5,
        typeInt16 = 6,
        typeInt32 = 7,
        typeInt64 = 8,
        typeFloat32 = 9,
        typeFloat64 = 10,
        typeString = 11,
        typeBytes = 12,
        typeBool = 13,
        typeProto = 14,
    }
}
