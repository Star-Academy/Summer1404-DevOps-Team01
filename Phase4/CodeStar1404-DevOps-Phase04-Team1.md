# HA Concept

Feel free to answer in Persian

- What is the importance of having a HA infrastructure?
  
  زیرساخت با دسترسی بالا (High Availability یا HA) تضمین می‌کند که سرویس‌ها و اپلیکیشن‌های شما همیشه در دسترس باشند و حتی در صورت بروز خطا یا خرابی، کاربران بتوانند به سرویس‌ها دسترسی داشته باشند. این موضوع برای کسب‌وکارها حیاتی است چون هرگونه قطعی می‌تواند باعث از دست رفتن درآمد، کاهش رضایت مشتری و آسیب به اعتبار شرکت شود.

- If we can afford 1s of downtime per day, what's our availabilty percantage for a year? explain.
  
  اگر بتوانیم روزانه فقط ۱ ثانیه قطعی داشته باشیم، درصد دسترسی سالانه را این‌گونه محاسبه می‌کنیم:
  - تعداد ثانیه‌های سال: ۳۶۵ × ۲۴ × ۶۰ × ۶۰ = ۳۱,۵۳۶,۰۰۰ ثانیه
  - مجموع قطعی سالانه: ۱ × ۳۶۵ = ۳۶۵ ثانیه
  - درصد دسترسی: ((۳۱,۵۳۶,۰۰۰ - ۳۶۵) / ۳۱,۵۳۶,۰۰۰) × ۱۰۰ ≈ ۹۹٫۹۹٪
  این یعنی زیرساخت شما تقریباً ۴ عدد ۹ (four nines) دسترسی دارد که بسیار عالی است.

- What is failover and how can it help us achive HA?
  
  Failover یعنی وقتی یک سرور یا سرویس دچار مشکل شد، به‌صورت خودکار درخواست‌ها به سرور یا سرویس جایگزین منتقل شوند تا سرویس‌دهی قطع نشود. این کار باعث می‌شود حتی در صورت خرابی بخشی از سیستم، کاربران همچنان بتوانند از سرویس استفاده کنند و HA حفظ شود.

- What are some methods to ensure HA and DR in our systems?

  استفاده از سرورهای متعدد و Load Balancer برای توزیع بار و جلوگیری از تک نقطه خرابی (Single Point of Failure)

  پیاده‌سازی مانیتورینگ و آلارم برای شناسایی سریع مشکلات و واکنش خودکار

  استفاده از دیتابیس‌های Replicated و Clustering برای حفظ داده‌ها و دسترسی بالا

  داشتن برنامه Disaster Recovery (DR) شامل بکاپ‌گیری منظم و تست بازیابی اطلاعات

- Explain chaos engineering and name some tools that help us with implementing it.
  
  مهندسی آشوب (Chaos Engineering) روشی است برای تست مقاومت و پایداری سیستم‌ها با ایجاد اختلالات عمدی (مثل قطع سرور یا شبکه) تا نقاط ضعف سیستم شناسایی و رفع شوند. ابزارهایی مثل Chaos Monkey (از Netflix)، Gremlin و LitmusChaos برای پیاده‌سازی این روش استفاده می‌شوند.

## Review

Link to your PR:  
[PR](https://github.com/Star-Academy/Summer1404-DevOps-Team01/pull/4)  

- [ ] Your PR is reviewed and approved by both mentors
- [ ] Your PR is merged
