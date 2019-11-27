using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

namespace Kongfu
{
    public class Joystick : ScrollRect
    {
        public event Action OnJoystickBegin;
        public event Action<float, float> OnJoystickDragging;
        public event Action OnJoystickEnd;

        public override void OnBeginDrag(PointerEventData eventData)
        {
            this.dragging = true;

            if (null != this.OnJoystickBegin)
            {
                this.OnJoystickBegin();
            }
        }

        public override void OnScroll(PointerEventData data)
        {
        }

        public override void OnDrag(PointerEventData eventData)
        {
            base.OnDrag(eventData);

            var contentPostion = this.content.anchoredPosition;
            if (contentPostion.magnitude > this.radius)
            {
                contentPostion = contentPostion.normalized * this.radius;
                SetContentAnchoredPosition(contentPostion);
            }
        }

        public override void OnEndDrag(PointerEventData eventData)
        {
            base.OnEndDrag(eventData);

            this.dragging = false;

            if (null != this.OnJoystickEnd)
            {
                this.OnJoystickEnd();
            }
        }

        protected override void Start()
        {
            this.inertia = false;
            this.movementType = MovementType.Unrestricted;

            //计算摇杆块的半径  
            this.radius = (transform as RectTransform).sizeDelta.x * 0.5f;
        }

        protected override void OnDisable()
        {
            this.content.localPosition = Vector3.zero;
        }

        private void Update()
        {
            this.offset = this.content.localPosition / this.radius;

            if (this.dragging)
            {
                if (null != OnJoystickDragging)
                {
                    this.OnJoystickDragging(this.offset.x, this.offset.y);
                }
            }
            else
            {
                float x = Mathf.Lerp(this.offset.x * this.radius, 0f, RecoveryTime);
                float y = Mathf.Lerp(this.offset.y * this.radius, 0f, RecoveryTime);
                this.content.localPosition = new Vector3(x, y, this.content.localPosition.z);
            }
        }

        private float radius = 0;
        private bool dragging = false;
        private Vector3 offset = Vector3.zero;

        private const float RecoveryTime = 0.1f;    //摇杆回退时间
    }
}
