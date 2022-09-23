package com.lab3;

import java.util.ArrayList;

import static java.lang.Thread.sleep;

public class Forest {

    Pot pot;
    ArrayList<Thread> bees;
    Bear bear;

    Forest() {

        pot = new Pot();

        bear = new Bear();
        new Thread(bear).start();

        bees = new ArrayList<>();
        for (int i = 0; i < 10; i++) {
            bees.add(new Thread(new Bee()));
            bees.get(i).start();
        }

    }


    class Bee implements Runnable{
        Bee(){ }

        public void run(){
            while (true) {
                pot.addHoney();
            }
        }
    }

    class Bear implements Runnable{

        Bear(){

        }
        public void run(){
            while (true) {
                pot.clearPot();
            }
        }
    }

    class Pot{
        private int value;
        private int max;
        Pot(){
            this.max = 10;
            value = 0;
        }

        public synchronized void addHoney(){
            while (value == this.max){
                try {
                    wait();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            value++;
            System.out.println("Bee added honey. 🐝 Pot value: " + value);
            notify();
        }

        public synchronized void clearPot(){
            while (value < this.max){
                try {
                    wait();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            value = 0;
            System.out.println("Bear ate all honey 🍯");
            notify();
        }
    }
}
