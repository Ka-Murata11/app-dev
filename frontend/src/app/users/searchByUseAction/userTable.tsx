"use client";

import { GetUsers } from "@/fetcher/user";
import { User } from "@/model/user";
import { useActionState, useEffect } from "react";
import Link from 'next/link'

export default function UserTable() {
    const [users, action] = useActionState(GetUsers, []);

    return (
        <>
            <h1>ユーザ管理</h1>
            <form action={action}>
                <label htmlFor="query">
                Search User:&nbsp;
                <input type="text" id="query" name="query" />
                </label>
                <button type="submit">Submit</button>
            </form>
            <ul>
                {users.map((user: User) => (
                    <li key={user.id}>
                        {user.name} - {user.email} - {user.role} - {user.job}
                    </li>
                ))}
            </ul>

            <div>
                <Link href='/users'>
                    <button>戻る</button>
                </Link>
            </div>
        </>
    )
}