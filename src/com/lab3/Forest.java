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
            synchronized (bear) {
                while (true) {
                    if (!pot.isFull()) {
                        pot.addHoney();
                        System.out.println("Bee added 1 honey");
                        try {
                            sleep(300);
                        } catch (InterruptedException e) {
                            e.printStackTrace();
                        }
                    } else {
                        bear.notify();
                    }
                }
            }
        }
    }

    class Bear implements Runnable{

        Bear(){

        }
        public void run(){
            while (true) {
                try {
                    wait();
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
                if (pot.isFull()) {
                    pot.clearPot();
                    System.out.println("Bear ate all honey");
                }
            }
        }
    }

    class Pot{
        private int pot;
        final int max_pot = 10;
        public Pot(){
            clearPot();
        }
        public boolean isFull(){
            return pot >= max_pot;
        }
        public void addHoney(){
            this.pot++;
        }
        public void clearPot(){
            this.pot = 0;
        }

    }
}
