"use client"; // Marca o arquivo como um componente cliente

import { useState, useEffect } from "react";
import { Menu, X } from "lucide-react";
import { useRouter } from "next/navigation"; // Importando hooks da Next Navigation
import ActivateLink from "@/components/linkActivate";
import Button from "@/components/button";
import Link from "next/link";

export default function Navbar() {
    const [isOpen, setIsOpen] = useState(false); // Controle do estado do menu
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [activeSection, setActiveSection] = useState("");
    const router = useRouter(); // Hook de navegação

    useEffect(() => {
        // Verifica se está no cliente antes de acessar o localStorage
        if (typeof window !== "undefined") {
            const token = sessionStorage.getItem("token");
            const username = sessionStorage.getItem("username");
            setIsAuthenticated(!!token && !!username);

            // Função para atualizar a seção ativa ao rolar a página
            const handleScroll = () => {
                const sections = document.querySelectorAll("section");
                let currentSection = "";

                sections.forEach((section) => {
                    const sectionTop = section.offsetTop;
                    const sectionHeight = section.offsetHeight;
                    if (window.scrollY >= sectionTop - sectionHeight / 3) {
                        currentSection = section.id;
                    }
                });

                setActiveSection(currentSection);
            };

            window.addEventListener("scroll", handleScroll);

            return () => {
                window.removeEventListener("scroll", handleScroll);
            };
        }
    }, []);

    const toggleMenu = () => setIsOpen(!isOpen);

    const handleLogout = () => {
        // Remove os dados de autenticação
        sessionStorage.removeItem("token");
        sessionStorage.removeItem("user_id");
        sessionStorage.removeItem("username");
        setIsAuthenticated(false);

        // Navegação após logout usando o Next.js Navigation
        router.push("/login"); // Navega para a página de login
    };

    return (
        <>
            <header>
                <nav className="w-full text-white p-4 fixed z-10 bg-mainColor">
                    <div className="container mx-auto flex items-center justify-between">
                        {/* Logo */}
                        <div className="logo flex justify-start items-center">
                            <Link href="/" className="font-archivo-black text-[20px] md:text-[30px]">
                                BARBERFAST
                            </Link>
                        </div>

                        {/* Links de Navegação (Desktop) */}
                        <div className="md:flex space-x-4 hidden justify-end px-5 items-center flex-grow">
                            <ActivateLink href="/#index" name="Home" activeSection={activeSection} />
                            <ActivateLink href="/#sobre" name="Sobre" activeSection={activeSection} />
                            <ActivateLink href="/#servicos" name="Serviços" activeSection={activeSection} />
                            <ActivateLink href="/#precos" name="Preços" activeSection={activeSection} />
                            <ActivateLink href="/#avaliacoes" name="Avaliações" activeSection={activeSection} />
                            <ActivateLink href="/#galeria" name="Galeria" activeSection={activeSection} />
                            <ActivateLink href="/#contato" name="Contato" activeSection={activeSection} />
                        </div>

                        {/* Menu Hamburguer (Aparece em todos os dispositivos) */}
                        <div className="flex items-center">
                            <button onClick={toggleMenu}>
                                {isOpen ? (
                                    <X className="w-6 h-6 text-white" />
                                ) : (
                                    <Menu className="w-6 h-6 text-white" />
                                )}
                            </button>
                        </div>
                    </div>

                    {/* Menu Deslizante (Aparece Quando Aberto) */}
                    <div
                        className={`absolute top-0 left-0 w-full bg-mainColor p-6 transition-transform duration-300 ease-in-out${
                            isOpen
                                ? "translate-y-0 opacity-100 visible z-20"
                                : "-translate-y-full opacity-0 invisible z-0"
                        }`}
                    >
                        {/* Fechar a aba clicando no X */}
                        <button
                            onClick={toggleMenu}
                            className="absolute top-4 right-4 z-30"
                        >
                            <X className="w-6 h-6 text-white" />
                        </button>

                        <div className="flex flex-col items-start space-y-4">
                            {/* Links do Menu */}
                            <ActivateLink href="/#index" name="Home" activeSection={activeSection} />
                            <ActivateLink href="/#sobre" name="Sobre" activeSection={activeSection} />
                            <ActivateLink href="/#servicos" name="Serviços" activeSection={activeSection} />
                            <ActivateLink href="/#precos" name="Preços" activeSection={activeSection} />
                            <ActivateLink href="/#avaliacoes" name="Avaliações" activeSection={activeSection} />
                            <ActivateLink href="/#galeria" name="Galeria" activeSection={activeSection} />
                            <ActivateLink href="/#contato" name="Contato" activeSection={activeSection} />

                            {isAuthenticated ? (
                                <>
                                    <Link href="/agendar" className="text-white hover:text-orange-500">
                                        Agendar
                                    </Link>
                                    <Link href="/agendamentos" className="text-white hover:text-orange-500">
                                        Meus Agendamentos
                                    </Link>
                                    <Link href="#" onClick={handleLogout} className="text-white hover:text-orange-500">
                                        <Button text="Sair" variant="primary" size="sm" />
                                    </Link>
                                </>
                            ) : (
                                <Link href="/login">
                                    <Button text="Login" variant="primary" size="sm" />
                                </Link>
                            )}
                        </div>
                    </div>
                </nav>
            </header>
        </>
    );
}
