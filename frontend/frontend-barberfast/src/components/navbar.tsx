import { useState, useEffect } from 'react';
import ActivateLink from "@/components/linkActivate";
import Button from './button';
import Link from 'next/link';

export default function Navbar() {
    const [activeSection, setActiveSection] = useState('');
    const [isClient, setIsClient] = useState(false);

    // Atualiza o estado isClient após a montagem no cliente
    useEffect(() => {
        setIsClient(true);
        
        const handleScroll = () => {
            const sections = document.querySelectorAll('section');
            let currentSection = '';

            sections.forEach((section) => {
                const sectionTop = section.offsetTop;
                const sectionHeight = section.offsetHeight;
                if (window.scrollY >= sectionTop - sectionHeight / 3) {
                    currentSection = section.id;
                }
            });

            setActiveSection(currentSection);
        };

        window.addEventListener('scroll', handleScroll);
        return () => {
            window.removeEventListener('scroll', handleScroll);
        };
    }, []); // A dependência vazia garante que o efeito seja executado apenas uma vez

    // Não renderize nada até que o código tenha rodado no cliente
    if (!isClient) return null;

    return (
        <>
            <header>
                <nav className="w-full text-white p-4 fixed z-10 bg-mainColor overflow-hidden">
                    <div className="container mx-auto flex items-center justify-between">
                        {/* Logo */}
                        <div className="logo flex justify-start items-center">
                            <a href="/" className="font-archivo-black text-[20px] md:text-[30px]">BARBERFAST</a>
                        </div>

                        {/* Links de Navegação */}
                        <div className="md:flex space-x-4 hidden justify-end px-5 items-center flex-grow">
                            <ActivateLink href="/#index" name="Home" activeSection={activeSection} />
                            <ActivateLink href="/#sobre" name="Sobre" activeSection={activeSection} />
                            <ActivateLink href="/#servicos" name="Serviços" activeSection={activeSection} />
                            <ActivateLink href="/#precos" name="Preços" activeSection={activeSection} />
                            <ActivateLink href="/#avaliacoes" name="Avaliações" activeSection={activeSection} />
                            <ActivateLink href="/#galeria" name="Galeria" activeSection={activeSection} />
                            <ActivateLink href="/#contato" name="Contato" activeSection={activeSection} />
                        </div>

                        {/* Botão de Login à Direita */}
                        <div className="login-nav ml-auto">
                            <Link href={`/login`}> 
                                <Button text="Login" variant="primary" size="sm"/>
                            </Link>
                        </div>
                    </div>
                </nav>
            </header>
        </>
    );
}
