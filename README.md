# url-shortener

using serverless framework and aws lambda to implement url shortener

## use case

- URL shortening: given a long URL => return a much shorter URL
- URL redirecting: given a shorter URL => redirect to the original URL
- High availability, scalability, and fault tolerance considerations

## Back of the envelope estimation

- Write operation: 100 million URLs are generated per day.
- Write operation per second: 100 million / 24 /3600 = 1160
- Read operation: Assuming ratio of read operation to write operation is 10:1, read operation per second: 1160 * 10 = 11,600
- Assuming the URL shortener service will run for 10 years, this means we must support 100 million * 365= 36.5 billion records.
- Assume average URL length is 100.
- Storage requirement over 10 years: 365 billion *100 bytes* 10 years = 365 TB

## design

### hash

### Base 62 conversion

## web/page

### pnpm

[pnpm node version,与node版本的兼容性](https://pnpm.io/installation#compatibility)

切换版本:
npm i pnpm@7.3.0 -g

npm i pnpm@8.3.0 -g

### nvm

https://github.com/jorgebucaran/nvm.fish

https://github.programnotes.cn/nvm-sh/nvm


### config

```json
//    "src": "/(.*)",
//    "dest": "/$1" 必須有$1,否則樣式不對 且控制台報錯: Uncaught SyntaxError: Unexpected token <
"routes": [
    {
        "src": "/_next/(.*)",
        "dest": "/_next/$1"
    },
    {
        "src": "/api/(.*)",
        "dest": "/api/short.go"
    },
    {
        "src": "/(.*)",
        "dest": "/$1"
    }
]
```

https://github.com/vercel/next.js/issues/7469