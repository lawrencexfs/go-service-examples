using UnityEngine;

namespace Kongfu
{
    public class CameraFollow : MonoBehaviour
    {
        public void SetTarget(GameObject target)
        {
            this.target = target;
        }

        private void Start()
        {
            transform.eulerAngles = CameraRotation;
        }

        private void Update()
        {
            if (null != this.target)
            {
                transform.position = this.target.transform.position + CameraOffset;
            }
        }

        private GameObject target = null;

        public readonly Vector3 CameraOffset = new Vector3(0f, 4.5f, -3f); //相机偏移
        private readonly Vector3 CameraRotation = new Vector3(55f, 0f, 0f); // 相机角度
    }
}