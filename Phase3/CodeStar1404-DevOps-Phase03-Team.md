Feel free to answer in Persian

- What is the difference between a static and a dynamic web server?

  **سرور استاتیک** فقط فایل‌های ثابت مانند HTML، CSS و تصاویر را بدون تغییر به کاربر ارائه می‌دهد. اما **سرور داینامیک** می‌تواند بر اساس درخواست کاربر، محتوای پویا تولید کند (مثلاً با اجرای کدهای سمت سرور مثل PHP یا Python).

- What is the difference between a load balancer and a reverse proxy?

  **لود بالانسر** ترافیک ورودی را بین چند سرور تقسیم می‌کند تا بار کاری متعادل شود. **ریورس پراکسی** درخواست‌های کاربران را دریافت کرده و به سرورهای داخلی منتقل می‌کند و معمولاً برای امنیت، کش و مدیریت ترافیک استفاده می‌شود. **لود بالانسر** می‌تواند نوعی **ریورس پراکسی** باشد.

- Name some load balancing strategies and explain how each of them works.

  - `Round Robin`: هر درخواست به ترتیب به سرورها ارسال می‌شود.
  - `Weighted Round Robin`: سرورها بر اساس وزن مشخص شده، تعداد بیشتری یا کمتری درخواست دریافت می‌کنند.
  - `Least Connections`: درخواست جدید به سروری می‌رود که کمترین تعداد اتصال فعال را دارد.
  - `IP Hash`: بر اساس هش آدرس آی پی کاربر، درخواست همیشه به یک سرور خاص ارسال می‌شود.
  - `Peak Exponentially Weighted Moving Average`: این روش به طور پویا وزن سرورها را بر اساس بار فعلی آنها تنظیم می‌کند.

- Explain some of the shortcomings of random load balancing strategy.

  در این روش ممکن است برخی سرورها بار بیشتری نسبت به بقیه دریافت کنند و تعادل بار به خوبی برقرار نشود. همچنین ممکن است اتصال‌های مرتبط به یک کاربر به سرورهای مختلفی برود که باعث کند شدن اتصال کاربر شود و یا با پر شدن صف سرورها، برخی درخواست‌ها رد شوند.

## Review

Link to your PR:  
[PR](https://github.com/Star-Academy/Summer1404-DevOps-Team01/pull/1)  

- [ ] Your PR is reviewed and approved by both mentors
- [ ] Your PR is merged
