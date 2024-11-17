'use client';

import { useRouter } from "next/router";

interface NavButtonProps {
    page: string;
}

const NavButton: React.FC<NavButtonProps> = ({ page }) => {
    const router = useRouter();

    const handleNavigate = () => {
        router.push(page);
    }

    return (
        <button onClick={handleNavigate}>ボタン</button>
    )
}

export default NavButton;