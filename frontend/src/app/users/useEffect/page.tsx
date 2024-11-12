import { User } from '@/model/user';
import UserTableEffect from './userTableEffect';
import CreateUserForm from './addUser';
import Link from 'next/link'
import { GetUsers } from './fetcher';

const Users = async () => {
    const users: User[] = await GetUsers();

    return (
        <>
            <CreateUserForm/>
            <UserTableEffect users={users}/>
            <div>
                <Link href='/users'>
                    <button>戻る</button>
                </Link>
            </div>
        </>
    )
}

export default Users;