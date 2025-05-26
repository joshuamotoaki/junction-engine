import { motion } from "framer-motion";
import { GITHUB_URL } from "~/utils/constants";

const Footer = () => {
    const footerLinks = {
        legal: [
            { name: "Terms of Service", href: "/tos" },
            { name: "Privacy Policy", href: "/privacy" },
        ],
    };

    return (
        <footer className="relative bg-gray-50 border-t border-gray-200">
            <div className="max-w-7xl mx-auto px-8 py-16">
                {/* Top section with logo */}
                <div className="flex flex-col md:flex-row justify-between items-start mb-12">
                    <div className="mb-8 md:mb-0">
                        <h3 className="text-3xl font-bold tracking-tighter">
                            TIGERJUNCTION
                        </h3>
                        <p className="text-sm text-gray-500 mt-2 max-w-xs">
                            Princeton academic planning,{" "}
                            <span className="font-telegraf-slanted font-bold">
                                amplified
                            </span>
                            .
                        </p>
                    </div>

                    {/* Newsletter signup */}
                    <div className="max-w-sm">
                        <p className="text-xs tracking-wider text-gray-500 mb-3">
                            STAY UPDATED
                        </p>
                        <form className="flex gap-2">
                            <input
                                type="email"
                                placeholder="Enter your email"
                                className="flex-1 px-4 py-2 text-sm bg-white border border-gray-300 focus:outline-none focus:border-gray-900 transition-colors"
                            />
                            <button
                                type="submit"
                                className="px-6 py-2 bg-gray-900 text-white text-sm hover:bg-gray-800 transition-colors"
                            >
                                →
                            </button>
                        </form>
                    </div>
                </div>

                {/* Links grid */}
                <div className="grid grid-cols-2 md:grid-cols-4 gap-8 mb-12">
                    <div>
                        <h4 className="text-xs tracking-wider text-gray-500 mb-4">
                            LEGAL
                        </h4>
                        <ul className="space-y-2">
                            {footerLinks.legal.map((link) => (
                                <li key={link.name}>
                                    <a
                                        href={link.href}
                                        className="text-sm text-gray-700 hover:text-gray-900 transition-colors"
                                    >
                                        {link.name}
                                    </a>
                                </li>
                            ))}
                        </ul>
                    </div>
                </div>

                {/* Bottom section */}
                <div className="flex flex-col md:flex-row justify-between items-center pt-8 border-t border-gray-200">
                    <p className="text-xs text-gray-500 mb-4 md:mb-0">
                        © 2023 TigerApps. All rights reserved.
                    </p>

                    {/* Social links */}
                    <div className="flex items-center gap-6">
                        <a
                            target="_blank"
                            referrerPolicy="no-referrer"
                            href={GITHUB_URL}
                            className="text-gray-500 hover:text-gray-900 transition-colors"
                            rel="noreferrer"
                        >
                            <span className="text-xs tracking-wider">
                                GITHUB
                            </span>
                        </a>
                        <a
                            target="_blank"
                            referrerPolicy="no-referrer"
                            href="https://tigerapps.org"
                            className="text-gray-500 hover:text-gray-900 transition-colors"
                            rel="noreferrer"
                        >
                            <span className="text-xs tracking-wider">
                                TIGERAPPS
                            </span>
                        </a>
                    </div>
                </div>

                {/* Decorative element */}
                <motion.div
                    className="absolute bottom-0 right-0 text-[200px] font-black text-gray-100/50 select-none pointer-events-none"
                    initial={{ opacity: 0 }}
                    whileInView={{ opacity: 1 }}
                    transition={{ duration: 1 }}
                >
                    TJ
                </motion.div>
            </div>
        </footer>
    );
};

export default Footer;
