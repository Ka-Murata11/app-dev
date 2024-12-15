import { User } from "@/model/user";
import { cookies } from "next/headers";

export async function GET() {
    const cookieStore = await cookies();
    const cookie = cookieStore.get('Authorization')?.value;
    const headers = new Headers();
    if (cookie) {
        headers.append('cookie', cookie);
    }
    const res = await fetch(`http://backend:1323/api/users`, {
        method: 'GET',
        headers,
    });
    const data = await res.json();
    // const resCookie = res.headers.get('set-cookie');
    // if (resCookie) {
    //     const cookieStore = await cookies();
    //     cookieStore.set({name: 'Authorization', value: resCookie});
    // }
    const users: User[] = data.users;
    return users;
}