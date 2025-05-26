import { useEffect, useRef, useState } from "react";
import {
    motion,
    useScroll,
    useTransform,
    useMotionValue,
    useSpring,
} from "motion/react";

const HeroSection = () => {
    const containerRef = useRef(null);
    const { scrollY } = useScroll();
    const [mousePosition, setMousePosition] = useState({ x: 0, y: 0 });

    // Parallax transforms
    const y1 = useTransform(scrollY, [0, 500], [0, -100]);
    const y2 = useTransform(scrollY, [0, 500], [0, -50]);
    const rotate = useTransform(scrollY, [0, 500], [0, -5]);

    // Mouse movement
    const mouseX = useMotionValue(0);
    const mouseY = useMotionValue(0);
    const springConfig = { damping: 15, stiffness: 150 };
    const xSpring = useSpring(mouseX, springConfig);
    const ySpring = useSpring(mouseY, springConfig);

    useEffect(() => {
        const handleMouseMove = (e) => {
            setMousePosition({ x: e.clientX, y: e.clientY });
            mouseX.set((e.clientX - window.innerWidth / 2) * 0.01);
            mouseY.set((e.clientY - window.innerHeight / 2) * 0.01);
        };

        window.addEventListener("mousemove", handleMouseMove);
        return () => window.removeEventListener("mousemove", handleMouseMove);
    }, [mouseX, mouseY]);

    return (
        <section
            ref={containerRef}
            className="relative min-h-screen overflow-hidden bg-white"
        >
            {/* Grain texture */}
            <div className="absolute inset-0 opacity-[0.04] mix-blend-multiply pointer-events-none">
                <svg width="100%" height="100%">
                    <filter id="noiseFilter">
                        <feTurbulence
                            type="fractalNoise"
                            baseFrequency="2"
                            numOctaves="1"
                        />
                    </filter>
                    <rect
                        width="100%"
                        height="100%"
                        filter="url(#noiseFilter)"
                    />
                </svg>
            </div>

            {/* Dynamic grid lines */}
            <div className="absolute inset-0 pointer-events-none">
                <motion.div
                    className="absolute top-0 left-[15%] w-[1px] h-full bg-gray-200"
                    style={{ scaleY: useTransform(scrollY, [0, 300], [0, 1]) }}
                />
                <motion.div
                    className="absolute top-[30%] left-0 w-full h-[1px] bg-gray-200"
                    style={{ scaleX: useTransform(scrollY, [0, 300], [0, 1]) }}
                />
            </div>

            {/* Floating elements */}
            <motion.div
                className="absolute top-[10%] right-[5%] text-xs tracking-[0.3em] text-gray-400"
                style={{ x: xSpring, y: ySpring }}
                initial={{ opacity: 0, x: 50 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: 1.2 }}
            >
                EST. 2023
            </motion.div>

            <motion.div
                className="absolute bottom-[15%] left-[5%] text-xs text-gray-400 -rotate-90 origin-left"
                style={{ y: y2 }}
            >
                SCROLL TO EXPLORE
            </motion.div>

            {/* Main typography */}
            <div className="relative h-screen">
                {/* TIGER - Large, bold, off-center */}
                <motion.div
                    className="absolute top-[15%] left-[8%]"
                    style={{ y: y1 }}
                >
                    <motion.h1
                        className="text-[12vw] font-black leading-[0.8] tracking-tighter"
                        initial={{
                            clipPath: "polygon(0 0, 0 0, 0 100%, 0 100%)",
                        }}
                        animate={{
                            clipPath: "polygon(0 0, 100% 0, 100% 100%, 0 100%)",
                        }}
                        transition={{ duration: 1, ease: [0.77, 0, 0.175, 1] }}
                    >
                        <span className="block text-gray-900">TIGER</span>
                    </motion.h1>
                </motion.div>

                {/* JUNCTION - Rotated, different position */}
                <motion.div
                    className="absolute top-[35%] right-[10%]"
                    style={{ rotate, x: mouseX }}
                >
                    <motion.h1
                        className="text-[10vw] font-light leading-none mix-blend-difference"
                        initial={{ opacity: 0, y: 50 }}
                        animate={{ opacity: 1, y: 0 }}
                        transition={{ duration: 0.8, delay: 0.3 }}
                    >
                        <span className="block text-gray-900">JUNCTION</span>
                    </motion.h1>
                </motion.div>

                {/* Experimental text block */}
                <motion.div
                    className="absolute top-[60%] left-[20%] max-w-sm"
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1 }}
                    transition={{ delay: 0.8 }}
                >
                    <p className="text-sm leading-relaxed text-gray-600">
                        Princeton&apos;s{" "}
                        <span className="font-bold font-telegraf-slanted">
                            premier
                        </span>{" "}
                        academic planning platform, completely{" "}
                        <span className="font-telegraf-slanted">
                            reimagined
                        </span>
                        . Transform how you navigate your academic journey.
                    </p>
                    <div className="mt-4 flex items-center gap-2">
                        <span className="text-[10px] tracking-wider text-gray-400">
                            01
                        </span>
                        <div className="h-[1px] w-20 bg-gray-300" />
                        <span className="text-[10px] tracking-wider text-gray-400">
                            EXPLORE
                        </span>
                    </div>
                </motion.div>

                {/* Scattered UI elements */}
                <motion.div
                    className="absolute top-[25%] left-[50%] -translate-x-1/2"
                    style={{ y: y2 }}
                    whileHover={{ scale: 1.1 }}
                >
                    <div className="w-32 h-32 border border-gray-300 rounded-full flex items-center justify-center">
                        <span className="text-xs tracking-wider text-gray-500">
                            ENTER
                        </span>
                    </div>
                </motion.div>

                {/* Navigation style CTA */}
                <motion.nav
                    className="absolute bottom-[10%] right-[10%] flex flex-col items-end gap-2"
                    initial={{ opacity: 0, x: 50 }}
                    animate={{ opacity: 1, x: 0 }}
                    transition={{ delay: 1 }}
                >
                    <motion.button
                        className="group relative overflow-hidden"
                        whileHover={{ x: -5 }}
                    >
                        <span className="relative z-10 text-sm tracking-wider text-gray-900 font-medium">
                            LOG IN →
                        </span>
                        <motion.div
                            className="absolute inset-0 bg-gray-900"
                            initial={{ x: "-100%" }}
                            whileHover={{ x: 0 }}
                            transition={{ duration: 0.3 }}
                        />
                        <motion.span
                            className="absolute inset-0 flex items-center text-sm tracking-wider text-white font-medium opacity-0 group-hover:opacity-100"
                            transition={{ duration: 0.3 }}
                        >
                            LOG IN →
                        </motion.span>
                    </motion.button>

                    <div className="text-xs text-gray-400 text-right">
                        <div>ACCOUNT</div>
                        <div>ACCESS</div>
                    </div>
                </motion.nav>

                {/* Abstract shapes */}
                <motion.div
                    className="absolute top-[45%] left-[60%] w-px h-40 bg-gray-300 origin-top"
                    style={{ rotate: useTransform(scrollY, [0, 300], [0, 45]) }}
                />

                <motion.div
                    className="absolute top-[70%] right-[30%] w-20 h-20 border border-gray-300"
                    style={{
                        rotate: useTransform(scrollY, [0, 300], [0, -90]),
                    }}
                />

                {/* Mouse follower */}
                <motion.div
                    className="fixed w-4 h-4 bg-gray-900 rounded-full mix-blend-difference pointer-events-none z-50"
                    style={{
                        left: mousePosition.x - 8,
                        top: mousePosition.y - 8,
                    }}
                />
            </div>

            {/* Side text */}
            <div className="absolute top-0 right-0 h-full w-12 flex items-center justify-center">
                <div className="transform rotate-90 text-xs tracking-[0.3em] text-gray-400">
                    TIGERAPPS © 2023
                </div>
            </div>
        </section>
    );
};

export default HeroSection;
