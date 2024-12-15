"use server"

import { cookies } from 'next/headers';

export async function Login(
    data: FormData,
) {
    console.log("server")
    const { user_id, password } = Object.fromEntries(data.entries());
    const res = await fetch(`http://backend:1323/signin`, {
        method:'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({user_id, password}),
        credentials: 'include',
    });
        // レスポンスヘッダーをログに出力
        const headers = Array.from(res.headers.entries());
        console.log('Response Headers:', headers);
        if (res.ok) {
            const setCookieHeader = res.headers.get('Set-Cookie');
            if (setCookieHeader) {
                const cookieStore = await cookies();
                cookieStore.set({name: 'Authorization', value: setCookieHeader});
                console.log('login success');
            }
        } else {
            console.log('login failed');
        }
    // if (res.status === 200) {
    //     console.log('login success')
    // } else {
    //     console.log('login failed')
    // }
}