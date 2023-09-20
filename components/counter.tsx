'use client'

import { json } from 'node:stream/consumers'
import React, { ChangeEvent, useState } from 'react'

type ShortUrl = {
    url: string
    short: string
    target: string
}

type GetShortUrl = {
    code: number
    message: string
    data: ShortUrl
}

export default function Counter() {
    const [text, setText] = useState("")
    const [shortUrl, setShortUrl] = useState("")
    const [target, setTarget] = useState("")

    const handlernput = (event: ChangeEvent<HTMLInputElement>) => {
        setText(event.target.value)
    }
    const handleShort = () => {
        fetch("/api/url?url=" + text).
            then((res) => res.json()).
            then((data) => {
                let v: GetShortUrl = data as GetShortUrl
                setShortUrl(v.data.short)
                setTarget(v.data.target)
                return
            }).
            catch((err) => {
                console.log(err)
            })
    }
    const handleLong = () => {
        fetch("/api/short?url=" + text).
            then((res) => res.json()).
            then((data) => {
                let v: GetShortUrl = data as GetShortUrl
                setShortUrl("")
                setTarget(v.data.target)
                return
            }).
            catch((err) => {
                console.log(err)
            })
    }
    return (
        <div>
            <p className='input w-full max-w-xs text-center text-3xl'>url shortener</p>
            <div>
                <form>
                    <input type="text" placeholder="Type here"
                        className="input input-bordered p-3 input-md w-full max-w-xs text-base mt-3"
                        value={text} onChange={handlernput} />
                </form>
                <p className='input input-bordered w-full max-w-xs p-3  text-sm mt-2 mr-0 pr-0 '>{shortUrl}</p>
                <p className='input input-bordered w-full max-w-xs p-3  text-sm mt-2 mr-0 pr-0 '>{target}</p>
            </div>
            <div>
                <button className="btn w-32 rounded-full  mt-4 mr-16"
                    onClick={handleLong}>short</button>
                <button className="btn w-32 rounded-full  mt-3"
                    onClick={handleShort}>long</button>
            </div>
        </div>
    )
}