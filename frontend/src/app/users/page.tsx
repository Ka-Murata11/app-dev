import React from 'react';
import Link from 'next/link'

const Users = async () => {

    return (
        <>
            <h1>ユーザ管理 選択</h1>
            
            <div>
                <Link href='/users/searchByUseAction'>
                    <button>searchByUseAction</button>
                </Link>
            </div>
            <div>
                <Link href='/users/useEffect'>
                    <button>useEffect</button>
                </Link>
            </div>
            <div>
                <Link href='/'>
                    <button>戻る</button>
                </Link>
            </div>
        </>
    )
}

export default Users;