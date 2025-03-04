import Navbar from "@/components/navbar";
import "@/styles/globals.css";
import { useState } from "react";
import { useRouter } from 'next/router';

export default function Login() {
    // Estado para armazenar os dados do formulário
    const [formData, setFormData] = useState({
        email: "",
        senha: "",
    });

    // Estado para armazenar mensagens de erro
    const [errors, setErrors] = useState<{ [key: string]: string }>({});
    const [isSubmitting, setIsSubmitting] = useState(false);
    const router = useRouter();

    // Manipula mudanças nos campos do formulário
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData({ 
            ...formData, 
            [name]: value 
        });
    };

    // Manipula o envio do formulário
    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        let validationErrors: { [key: string]: string } = {};
    
        // Validação do email e senha
        if (!formData.email.includes("@")) validationErrors.email = "Email inválido.";
        if (!formData.senha) validationErrors.senha = "A senha é obrigatória.";
    
        setErrors(validationErrors);
    
        // Se houver erros, interrompe o envio
        if (Object.keys(validationErrors).length > 0) return;

        setIsSubmitting(true);

        try {
            // Envia os dados para a API
            const response = await fetch("http://192.168.1.4:5050/v1/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(formData),
            });
    
            const data = await response.json();
    
            if (response.ok) {
                // Armazena o token no localStorage
                sessionStorage.setItem("token", data.token);;
                sessionStorage.setItem("username", data.username);
                sessionStorage.setItem("user_id", data.user_id)

                // Redireciona para a página inicial em caso de sucesso
                router.push("/agendar");
            } else {
                // Define erro de login inválido
                setErrors({ email: "Email ou senha inválidos." });
                setIsSubmitting(false);
            }
        } catch (error) {
            console.error("Erro ao enviar dados:", error);
            setIsSubmitting(false);
        }
    };

    return (
        <>
            <Navbar />
            <section id="login" className="bg-login bg-cover min-h-screen flex flex-col justify-center items-center">
                <div className="overlay flex flex-col justify-center items-center bg-mainColor/75 absolute w-full h-full">
                    <div className="login flex flex-col md:flex-row items-center justify-center gap-8 p-6 md:p-12">
                        {/* Seção esquerda com texto explicativo */}
                        <div className="box-esquerda flex-1 text-center md:text-left px-4 md:px-[10%]">
                            <h1 className="text-xl text-center md:text-3xl text-white mb-6">
                                Faça seu <span className="text-laranja">LOGIN</span> para acessar sua conta e aproveitar todos os benefícios.
                            </h1>
                        </div>
                        {/* Seção direita com o formulário de login */}
                        <div className="box-direita flex-1 flex items-center justify-center p-6 rounded-xl w-full max-w-md">
                            <form onSubmit={handleSubmit} className="w-full font-archivo">
                                <h2 className="text-white text-2xl font-archivo-black mb-6 text-center">LOGIN</h2>
                                <div className="mb-4">
                                    {/* Campo de email */}
                                    <input  
                                        type="email" 
                                        name="email" 
                                        value={formData.email} 
                                        placeholder="Email"
                                        onChange={handleChange} 
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    />
                                    {errors.email && <p className="text-red-500 text-sm">{errors.email}</p>}
                                </div>
                                <div className="mb-4">
                                    {/* Campo de senha */}
                                    <input 
                                        type="password" 
                                        name="senha" 
                                        value={formData.senha}
                                        placeholder="Senha" 
                                        onChange={handleChange} 
                                        className="w-full px-4 py-2 rounded-full bg-white/10 text-white focus:outline-none focus:ring-2 focus:ring-laranja"
                                    />
                                    {errors.senha && <p className="text-red-500 text-sm">{errors.senha}</p>}
                                </div>
                                {/* Botão de envio */}
                                <button 
                                    type="submit" 
                                    className="w-full bg-laranja font-archivo-black text-white py-3 rounded-full font-bold text-lg hover:bg-orange-500 transition"
                                    disabled={isSubmitting}
                                >
                                    {isSubmitting ? 'Entrando...' : 'ENTRAR'}
                                </button>
                            </form>
                        </div>
                    </div>
                </div>
            </section>
        </>
    );
}
