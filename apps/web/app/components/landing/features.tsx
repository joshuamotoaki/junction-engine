import React, { useRef } from "react";
import { motion, useScroll, useTransform } from "framer-motion";

const FeaturesSection = () => {
    const containerRef = useRef(null);
    const { scrollYProgress } = useScroll({
        target: containerRef,
        offset: ["start end", "end start"],
    });

    const features = [
        {
            id: "01",
            title: "NEURAL SYNC",
            description:
                "Real-time collaborative intelligence that adapts to your workflow, learning from patterns and anticipating needs before they arise",
            tag: "INTELLIGENCE",
        },
        {
            id: "02",
            title: "QUANTUM MESH",
            description:
                "Distributed processing across infinite computational nodes, creating a seamless network that scales beyond traditional limitations",
            tag: "INFRASTRUCTURE",
        },
        {
            id: "03",
            title: "VOID PROTOCOL",
            description:
                "Zero-latency communication through dimensional barriers, enabling instant synchronization across any distance or platform",
            tag: "CONNECTIVITY",
        },
        {
            id: "04",
            title: "FLUX ENGINE",
            description:
                "Dynamic resource allocation with predictive scaling that automatically optimizes performance based on real-time demand patterns",
            tag: "PERFORMANCE",
        },
    ];

    // Create transforms for each dot with more spaced out progress
    const dotScales = [
        useTransform(scrollYProgress, [0, 0.2], [0.5, 1.5]),
        useTransform(scrollYProgress, [0.2, 0.4], [0.5, 1.5]),
        useTransform(scrollYProgress, [0.4, 0.6], [0.5, 1.5]),
        useTransform(scrollYProgress, [0.6, 0.8], [0.5, 1.5]),
    ];

    const dotColors = [
        useTransform(scrollYProgress, [0, 0.2], ["#d1d5db", "#1f2937"]),
        useTransform(scrollYProgress, [0.2, 0.4], ["#d1d5db", "#1f2937"]),
        useTransform(scrollYProgress, [0.4, 0.6], ["#d1d5db", "#1f2937"]),
        useTransform(scrollYProgress, [0.6, 0.8], ["#d1d5db", "#1f2937"]),
    ];

    return (
        <section ref={containerRef} className="relative min-h-[500vh] bg-white">
            {" "}
            {/* Much larger container */}
            {/* Grain overlay */}
            <div className="fixed inset-0 opacity-[0.03] mix-blend-multiply pointer-events-none z-10">
                <svg width="100%" height="100%">
                    <filter id="featureNoise">
                        <feTurbulence
                            type="fractalNoise"
                            baseFrequency="3"
                            numOctaves="1"
                        />
                    </filter>
                    <rect
                        width="100%"
                        height="100%"
                        filter="url(#featureNoise)"
                    />
                </svg>
            </div>
            <div className="sticky top-0 h-screen overflow-hidden">
                {/* Background elements */}
                <motion.div
                    className="absolute inset-0"
                    style={{
                        scale: useTransform(scrollYProgress, [0, 1], [1.5, 1]),
                        opacity: useTransform(
                            scrollYProgress,
                            [0, 0.3],
                            [0, 1]
                        ),
                    }}
                >
                    <div className="absolute top-[10%] left-[5%] w-[1px] h-[80%] bg-gray-200" />
                    <div className="absolute top-[50%] left-0 w-full h-[1px] bg-gray-100" />
                    <div className="absolute top-[20%] right-[15%] w-[1px] h-[60%] bg-gray-200" />
                </motion.div>

                {/* Section title */}
                <motion.div
                    className="absolute top-[5%] left-[5%]"
                    style={{
                        x: useTransform(scrollYProgress, [0, 0.2], [-100, 0]),
                        opacity: useTransform(
                            scrollYProgress,
                            [0, 0.2],
                            [0, 1]
                        ),
                    }}
                >
                    <h2 className="text-[8vw] font-black leading-none tracking-tighter text-gray-900">
                        FEATURES
                    </h2>
                    <div className="mt-2 flex items-center gap-2">
                        <span className="text-[10px] tracking-wider text-gray-400">
                            SECTION
                        </span>
                        <div className="h-[1px] w-20 bg-gray-300" />
                        <span className="text-[10px] tracking-wider text-gray-400">
                            02
                        </span>
                    </div>
                </motion.div>

                {/* Feature 1 - Gets 0-0.25 scroll progress */}
                <motion.div
                    className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 max-w-2xl w-full px-8"
                    style={{
                        y: useTransform(
                            scrollYProgress,
                            [0, 0.25],
                            [100, -100]
                        ),
                        opacity: useTransform(
                            scrollYProgress,
                            [0, 0.1, 0.2, 0.25],
                            [0, 1, 1, 0]
                        ),
                        scale: useTransform(
                            scrollYProgress,
                            [0, 0.1, 0.2, 0.25],
                            [0.8, 1, 1, 0.8]
                        ),
                    }}
                >
                    <span className="text-[300px] font-black text-gray-100 absolute -top-32 -left-16 -z-10">
                        01
                    </span>
                    <h3 className="text-6xl font-bold mb-6 leading-tight">
                        {features[0].title}
                    </h3>
                    <p className="text-gray-600 leading-relaxed text-xl mb-8">
                        {features[0].description}
                    </p>
                    <span className="inline-block text-sm tracking-[0.2em] text-gray-400 border border-gray-300 px-6 py-3">
                        {features[0].tag}
                    </span>
                </motion.div>

                {/* Feature 2 - Gets 0.25-0.5 scroll progress */}
                <motion.div
                    className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 max-w-2xl w-full px-8"
                    style={{
                        rotate: useTransform(
                            scrollYProgress,
                            [0.25, 0.5],
                            [15, 0]
                        ),
                        x: useTransform(scrollYProgress, [0.25, 0.5], [200, 0]),
                        opacity: useTransform(
                            scrollYProgress,
                            [0.25, 0.3, 0.45, 0.5],
                            [0, 1, 1, 0]
                        ),
                        scale: useTransform(
                            scrollYProgress,
                            [0.25, 0.3, 0.45, 0.5],
                            [0.8, 1, 1, 0.8]
                        ),
                    }}
                >
                    <div className="relative bg-white p-12 shadow-xl">
                        <span className="text-[200px] font-light text-gray-300 absolute -top-24 -right-12">
                            02
                        </span>
                        <h3 className="text-5xl font-medium mb-6">
                            {features[1].title}
                        </h3>
                        <p className="text-gray-500 leading-relaxed text-lg mb-8">
                            {features[1].description}
                        </p>
                        <span className="inline-block text-sm tracking-[0.2em] text-gray-400 border border-gray-300 px-6 py-3">
                            {features[1].tag}
                        </span>
                    </div>
                </motion.div>

                {/* Feature 3 - Gets 0.5-0.75 scroll progress */}
                <motion.div
                    className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 max-w-2xl w-full px-8"
                    style={{
                        scale: useTransform(
                            scrollYProgress,
                            [0.5, 0.6, 0.7, 0.75],
                            [0.6, 1, 1, 0.6]
                        ),
                        opacity: useTransform(
                            scrollYProgress,
                            [0.5, 0.55, 0.7, 0.75],
                            [0, 1, 1, 0]
                        ),
                        x: useTransform(
                            scrollYProgress,
                            [0.5, 0.75],
                            [-200, 0]
                        ),
                    }}
                >
                    <div className="flex items-center gap-12">
                        <span className="text-[280px] font-black text-gray-200 leading-none">
                            03
                        </span>
                        <div>
                            <h3 className="text-4xl font-bold tracking-tight mb-6">
                                {features[2].title}
                            </h3>
                            <p className="text-gray-500 leading-relaxed text-lg mb-8">
                                {features[2].description}
                            </p>
                            <span className="inline-block text-sm tracking-[0.2em] text-gray-400 border border-gray-300 px-6 py-3">
                                {features[2].tag}
                            </span>
                        </div>
                    </div>
                </motion.div>

                {/* Feature 4 - Gets 0.75-1 scroll progress */}
                <motion.div
                    className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 max-w-3xl w-full px-8"
                    style={{
                        y: useTransform(scrollYProgress, [0.75, 1], [100, -50]),
                        rotate: useTransform(
                            scrollYProgress,
                            [0.75, 1],
                            [-10, 5]
                        ),
                        opacity: useTransform(
                            scrollYProgress,
                            [0.75, 0.8, 0.95, 1],
                            [0, 1, 1, 0]
                        ),
                        scale: useTransform(
                            scrollYProgress,
                            [0.75, 0.8, 0.95, 1],
                            [0.8, 1, 1, 0.9]
                        ),
                    }}
                >
                    <div className="bg-gray-50 p-16 shadow-2xl relative">
                        <span className="text-[120px] tracking-[0.3em] text-gray-400 absolute top-8 left-8">
                            04
                        </span>
                        <div className="pt-20">
                            <h3 className="text-6xl font-light mb-8">
                                {features[3].title}
                            </h3>
                            <div className="w-full h-[2px] bg-gray-300 my-8" />
                            <p className="text-gray-600 leading-relaxed text-xl mb-10">
                                {features[3].description}
                            </p>
                            <span className="inline-block text-sm tracking-[0.2em] text-gray-400 border border-gray-300 px-6 py-3">
                                {features[3].tag}
                            </span>
                        </div>
                        <motion.div
                            className="absolute -bottom-8 -right-8 w-32 h-32 border-2 border-gray-300 rounded-full"
                            animate={{ rotate: 360 }}
                            transition={{
                                duration: 20,
                                repeat: Infinity,
                                ease: "linear",
                            }}
                        />
                    </div>
                </motion.div>

                {/* Floating elements */}
                <motion.div
                    className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-[400px] font-black text-gray-50 -z-20"
                    style={{
                        x: useTransform(scrollYProgress, [0, 1], [300, -300]),
                        y: useTransform(scrollYProgress, [0, 1], [200, -200]),
                        rotate: useTransform(scrollYProgress, [0, 1], [0, 360]),
                        opacity: useTransform(
                            scrollYProgress,
                            [0, 0.2, 0.8, 1],
                            [0.3, 0.1, 0.1, 0.3]
                        ),
                    }}
                >
                    ✦
                </motion.div>

                {/* Side navigation dots */}
                <div className="absolute right-8 top-1/2 -translate-y-1/2 flex flex-col gap-4">
                    {features.map((_, i) => (
                        <motion.div
                            key={i}
                            className="w-4 h-4 rounded-full bg-gray-300"
                            style={{
                                scale: dotScales[i],
                                backgroundColor: dotColors[i],
                            }}
                        />
                    ))}
                </div>
            </div>
        </section>
    );
};

export default FeaturesSection;
