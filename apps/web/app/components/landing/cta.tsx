import React, { useRef } from "react";
import { motion, useScroll, useTransform, useInView } from "framer-motion";

const CTASection = () => {
    const containerRef = useRef(null);
    const isInView = useInView(containerRef, { once: false, margin: "-100px" });
    const { scrollYProgress } = useScroll({
        target: containerRef,
        offset: ["start end", "end start"],
    });

    const scale = useTransform(scrollYProgress, [0, 0.5], [0.8, 1]);
    const rotate = useTransform(scrollYProgress, [0, 1], [-5, 5]);

    return (
        <section
            ref={containerRef}
            className="relative min-h-screen flex items-center justify-center overflow-hidden bg-white"
        >
            {/* Background texture */}
            <div className="absolute inset-0 opacity-[0.02]">
                <svg width="100%" height="100%">
                    <pattern
                        id="dots"
                        x="0"
                        y="0"
                        width="20"
                        height="20"
                        patternUnits="userSpaceOnUse"
                    >
                        <circle cx="2" cy="2" r="1" fill="#000" opacity="0.1" />
                    </pattern>
                    <rect width="100%" height="100%" fill="url(#dots)" />
                </svg>
            </div>

            {/* Floating background text */}
            <motion.div
                className="absolute inset-0 flex items-center justify-center -z-10"
                style={{ scale, rotate }}
            >
                <h3 className="text-[20vw] font-black text-gray-50 leading-none select-none">
                    JOIN
                </h3>
            </motion.div>

            {/* Main CTA content */}
            <div className="relative z-10 text-center max-w-4xl mx-auto px-8">
                <motion.div
                    initial={{ opacity: 0, y: 50 }}
                    animate={isInView ? { opacity: 1, y: 0 } : {}}
                    transition={{ duration: 0.8 }}
                >
                    {/* Small label */}
                    <motion.div
                        className="flex items-center justify-center gap-2 mb-8"
                        initial={{ width: 0 }}
                        animate={isInView ? { width: "auto" } : {}}
                        transition={{ duration: 0.6, delay: 0.2 }}
                    >
                        <div className="h-[1px] w-12 bg-gray-400" />
                        <span className="text-xs tracking-[0.3em] text-gray-500">
                            READY TO START?
                        </span>
                        <div className="h-[1px] w-12 bg-gray-400" />
                    </motion.div>

                    {/* Main headline */}
                    <h2 className="text-6xl md:text-8xl font-bold leading-[0.9] tracking-tighter mb-6">
                        <motion.span
                            className="block"
                            initial={{ clipPath: "inset(0 100% 0 0)" }}
                            animate={
                                isInView ? { clipPath: "inset(0 0% 0 0)" } : {}
                            }
                            transition={{
                                duration: 0.8,
                                ease: [0.77, 0, 0.175, 1],
                            }}
                        >
                            BUILD THE
                        </motion.span>
                        <motion.span
                            className="block text-gray-400"
                            initial={{ clipPath: "inset(0 100% 0 0)" }}
                            animate={
                                isInView ? { clipPath: "inset(0 0% 0 0)" } : {}
                            }
                            transition={{
                                duration: 0.8,
                                delay: 0.1,
                                ease: [0.77, 0, 0.175, 1],
                            }}
                        >
                            EXTRAORDINARY
                        </motion.span>
                    </h2>

                    <motion.p
                        className="text-xl text-gray-600 mb-12 max-w-2xl mx-auto"
                        initial={{ opacity: 0 }}
                        animate={isInView ? { opacity: 1 } : {}}
                        transition={{ duration: 0.8, delay: 0.4 }}
                    >
                        Join thousands of innovators already using TigerJunction
                        to transform their ideas into reality.
                    </motion.p>

                    {/* CTA Button - Large and experimental */}
                    <motion.div
                        initial={{ scale: 0 }}
                        animate={isInView ? { scale: 1 } : {}}
                        transition={{
                            duration: 0.5,
                            delay: 0.6,
                            type: "spring",
                        }}
                    >
                        <motion.button
                            className="group relative px-16 py-8 text-2xl font-medium tracking-wider bg-gray-900 text-white overflow-hidden"
                            whileHover={{ scale: 1.05 }}
                            whileTap={{ scale: 0.95 }}
                        >
                            <span className="relative z-10">LOG IN</span>

                            {/* Hover effect */}
                            <motion.div
                                className="absolute inset-0 bg-white"
                                initial={{ y: "100%" }}
                                whileHover={{ y: 0 }}
                                transition={{
                                    duration: 0.3,
                                    ease: "easeInOut",
                                }}
                            />
                            <motion.span
                                className="absolute inset-0 flex items-center justify-center text-gray-900 font-bold"
                                initial={{ opacity: 0 }}
                                whileHover={{ opacity: 1 }}
                                transition={{ duration: 0.3 }}
                            >
                                LOG IN →
                            </motion.span>

                            {/* Corner details */}
                            <div className="absolute top-0 left-0 w-8 h-8 border-t-2 border-l-2 border-white -translate-x-1 -translate-y-1" />
                            <div className="absolute bottom-0 right-0 w-8 h-8 border-b-2 border-r-2 border-white translate-x-1 translate-y-1" />
                        </motion.button>
                    </motion.div>

                    {/* Alternative action */}
                    <motion.p
                        className="mt-8 text-sm text-gray-500"
                        initial={{ opacity: 0 }}
                        animate={isInView ? { opacity: 1 } : {}}
                        transition={{ duration: 0.8, delay: 0.8 }}
                    >
                        No account yet?{" "}
                        <a
                            href="#"
                            className="underline hover:text-gray-900 transition-colors"
                        >
                            Contact Sales
                        </a>
                    </motion.p>
                </motion.div>

                {/* Side decorative elements */}
                <motion.div
                    className="absolute -left-20 top-1/2 -translate-y-1/2 text-8xl font-black text-gray-100 -rotate-90"
                    style={{
                        x: useTransform(scrollYProgress, [0, 1], [-50, 50]),
                    }}
                >
                    2025
                </motion.div>

                <motion.div
                    className="absolute -right-20 top-1/2 -translate-y-1/2 text-8xl font-black text-gray-100 rotate-90"
                    style={{
                        x: useTransform(scrollYProgress, [0, 1], [50, -50]),
                    }}
                >
                    BETA
                </motion.div>
            </div>

            {/* Floating shapes */}
            <motion.div
                className="absolute top-20 right-1/4 w-32 h-32 border border-gray-200 rounded-full"
                animate={{
                    rotate: 360,
                    scale: [1, 1.1, 1],
                }}
                transition={{
                    rotate: { duration: 20, repeat: Infinity, ease: "linear" },
                    scale: { duration: 4, repeat: Infinity, ease: "easeInOut" },
                }}
            />
        </section>
    );
};

export default CTASection;
