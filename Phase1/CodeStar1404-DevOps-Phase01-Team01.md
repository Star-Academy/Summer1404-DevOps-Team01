Feel free to answer in Persian

- Name some of Linux popular distributions:
  - Debian
  - Red-Hat
  - SUSE
  - ...
- What is the /proc directory is used for?  
  یک فایل‌سیستم مجازی است که اطلاعات مربوط به وضعیت جاری سیستم، پردازش‌ها، سخت‌افزار و هسته سیستم را به صورت فایل و پوشه نمایش می‌دهد.
- Where are system-wide configuration files are stored in linux directory hierarchy?  
  /etc
- How linux shell commands are executed?  
  دستورات شل توسط مفسر شل خوانده و اجرا می‌شوند؛ ابتدا دستور تجزیه می‌شود، سپس اجرا به صورت یک فرایند جدید انجام می‌گیرد.
- Name some popular linux commands and their respective usage:  
  - ls: نمایش لیست فایل‌ها و پوشه‌ها
  - cd: تغییر دایرکتوری جاری
  - cp: کپی کردن فایل یا پوشه
  - mv: جابجایی یا تغییر نام فایل یا پوشه
  - rm: حذف فایل یا پوشه
  - cat: نمایش محتوای فایل
  - grep: جستجو در متن فایل‌ها
  - ...
- What does this command do?  
  :(){:|:&};:  
  این یک fork bomb است که به طور بازگشتی خودش را اجرا می‌کند و باعث مصرف بیش از حد منابع سیستم و از کار افتادن سیستم می‌شود.
- How new packages are installed in linux?  
  بسته به توزیع، با استفاده از ابزارهایی مانند apt، yum، dnf یا zypper بسته‌ها نصب می‌شوند.
- Given the following ls command output, if danny is only inside the group danny, what files can he read?  

```bash
    drwxr-xr-x  2 danny danny 4.0K May 25 23:06 .
    drwxr-x--- 25 danny danny 4.0K Jul  8 22:19 ..
    -rwxrwxrwx  1 danny admin    0 May 25 23:02 a
    ----r--r--  1 danny games    0 May 25 23:02 b
    -r--rw----  1 root  root     0 May 25 23:02 c
    -r--r-----  1 root  danny    0 May 25 23:02 d
```

a, b, d

- What is . file and what happens if its permissions are set to 000?  
  . نمایانگر دایرکتوری جاری است و اگر مجوزهای آن 000 باشد، هیچ کاربری حتی مالک نمی‌تواند به آن دسترسی داشته باشد.
- How can you exit Vim?  
  :q یا :wq یا ZZ
- Who is the murderer running around in terminal city? How did you find them?  
  پاسخ بستگی به حل معمای clmystery دارد و باید با دنبال کردن سرنخ‌ها و بررسی فایل‌ها و مصاحبه‌ها، قاتل را پیدا کنید.

## Review

Link to your PR:  
[LINK TO YOUR PR]  



- [ ] Your PR is reviewed and approved by both mentors
- [ ] Your PR is merged
