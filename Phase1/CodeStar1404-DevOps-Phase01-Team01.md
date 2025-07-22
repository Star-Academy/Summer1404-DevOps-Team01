Feel free to answer in Persian
- Name some of Linux popular distributions:
    - Debian
    - Red-Hat
    - SUSE
    - ...
- What is the /proc directory is used for?  
```  یک فایل‌سیستم مجازی  است که اطلاعات مربوط به وضعیت جاری سیستم، پردازش‌ها ، سخت‌افزار و کرنل (هسته سیستم) را به صورت فایل و پوشه```
- Where are system-wide configuration files are stored in linux directory hierarchy?  
  ```/etc```
- How linux shell commands are executed?  
```  درواقع هر دستور توسط مفسر شل خوانده و اجرا می‌شود و اسمش رو توی محتویات متغیر مسیر(پث) که ادرس فایل های اجرایی توش هست میگرده و اون فایل باینری رو اجرا میکنه```
- Name some popular linux commands and their respective usage:  
    - ls: نمایش لیست فایل‌ها و پوشه‌ها
    - cd: تغییر دایرکتوری جاری
    - grep: جستجو در متن فایل‌ها
    - cat: نمایش محتوای فایل
- What does this command do?  
  :(){:|:&};:  
```  این داره یک تابع تعریف میکنه با :() و بعد خودش رو پایپ میکنه به دو تا از خودش و میفرسته به بک با & پس یه تابع بی نهایت میسازه و منابع و مصرف میکنه تا سیستم کرش کنه```
- How new packages are installed in linux?  
  
```بسته به توزیع، با استفاده از ابزارهایی مانند apt (در دبیان/اوبونتو)، yum یا dnf (در رد هت/فدورا)، zypper (در SUSE) و ... بسته‌ها نصب می‌شوند```
- Given the following ls command output, if danny is only inside the group danny, what files can he read?  
    ```
    drwxr-xr-x  2 danny danny 4.0K May 25 23:06 .
    drwxr-x--- 25 danny danny 4.0K Jul  8 22:19 ..
    -rwxrwxrwx  1 danny admin    0 May 25 23:02 a
    ----r--r--  1 danny games    0 May 25 23:02 b
    -r--rw----  1 root  root     0 May 25 23:02 c
    -r--r-----  1 root  danny    0 May 25 23:02 d
    ```
    [a, d, ., ..]
- What is . file and what happens if its permissions are set to 000?  
  `این نماینده پوشه فعلیه و اگر دسترسیشو به 000 تغییر بدیم دیگه نمیتونیم به اون پوشه دسترسی داشته باشیم و اگه توش باشیم 
ls
دیگه نمیتونیم بزنیم توش یا خارج شیم نمیتونیم دوباره برگردیم بهش`

- How can you exit Vim?  
```  esc + :q or :q!(not save changes) or :wq(save changes) or ZZ```
- Who is the murderer running around in terminal city? How did you find them?  
  `Jeremy Bowers`
we just followed the clues and used some of the hints

## Review

Link to your PR:  
[PR](https://github.com/Star-Academy/Summer1404-DevOps-Team01/pull/1)  

- [ ] Your PR is reviewed and approved by both mentors
- [ ] Your PR is merged
