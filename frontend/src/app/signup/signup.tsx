"use server"

export async function Signup(
    data: FormData,
) {
    const { user_id, job, email, password } = Object.fromEntries(data.entries());
    const res = await fetch(`http://backend:1323/signup`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ user_id, job, email, password}),
    });

    if (res.ok) {
        console.log('signup success');
    } else {
        console.log('signup failed');
    }
}